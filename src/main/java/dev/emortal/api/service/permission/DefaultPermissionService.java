package dev.emortal.api.service.permission;

import dev.emortal.api.grpc.permission.PermissionProto;
import dev.emortal.api.grpc.permission.PermissionServiceGrpc;
import dev.emortal.api.model.permission.Role;
import dev.emortal.api.utils.internal.GrpcErrorHelper;
import io.grpc.Channel;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

/**
 * The default implementation of {@link PermissionService} that uses a blocking stub to communicate with the gRPC server.
 */
@ApiStatus.Internal
public final class DefaultPermissionService implements PermissionService {

    private final PermissionServiceGrpc.PermissionServiceBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultPermissionService(@NotNull Channel channel) {
        this.grpc = PermissionServiceGrpc.newBlockingStub(channel);
    }

    @Override
    public @NotNull List<Role> getAllRoles() {
        var request = PermissionProto.GetAllRolesRequest.getDefaultInstance();
        PermissionProto.GetAllRolesResponse response = this.grpc.getAllRoles(request);
        return response.getRolesList();
    }

    @Override
    public @NotNull PermissionProto.PlayerRolesResponse getPlayerRoles(@NotNull UUID playerId) {
        var request = PermissionProto.GetPlayerRolesRequest.newBuilder().setPlayerId(playerId.toString()).build();
        return this.grpc.getPlayerRoles(request);
    }

    @Override
    public @NotNull CreateRoleResult createRole(@NotNull String id, int priority, @Nullable String displayName) {
        var requestBuilder = PermissionProto.RoleCreateRequest.newBuilder().setId(id).setPriority(priority);
        if (displayName != null) requestBuilder.setDisplayName(displayName);
        var request = requestBuilder.build();

        PermissionProto.CreateRoleResponse response;
        try {
            response = this.grpc.createRole(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus().getCode() == Status.Code.ALREADY_EXISTS) return CreateRoleResult.Error.ROLE_ALREADY_EXISTS;
            throw exception;
        }

        return new CreateRoleResult.Success(response.getRole());
    }

    @Override
    public @NotNull UpdateRoleResult updateRole(@NotNull RoleUpdate roleUpdate) {
        var requestBuilder = PermissionProto.RoleUpdateRequest.newBuilder().setId(roleUpdate.id());

        if (roleUpdate.priority() != -1) requestBuilder.setPriority(roleUpdate.priority());
        if (roleUpdate.displayName() != null) requestBuilder.setDisplayName(roleUpdate.displayName());
        requestBuilder.addAllSetPermissions(roleUpdate.setPermissions());
        requestBuilder.addAllUnsetPermissions(roleUpdate.unsetPermissions());

        var request = requestBuilder.build();

        PermissionProto.UpdateRoleResponse response;
        try {
            response = this.grpc.updateRole(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus().getCode() == Status.Code.NOT_FOUND) return UpdateRoleResult.Error.ROLE_NOT_FOUND;
            throw exception;
        }

        return new UpdateRoleResult.Success(response.getRole());
    }

    @Override
    public @NotNull AddRoleToPlayerResult addRoleToPlayer(@NotNull UUID playerId, @NotNull String roleId) {
        var request = PermissionProto.AddRoleToPlayerRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setRoleId(roleId)
                .build();

        try {
            this.grpc.addRoleToPlayer(request);
            return AddRoleToPlayerResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PermissionProto.AddRoleToPlayerError.class);

            return switch (error.getErrorType()) {
                case PLAYER_NOT_FOUND -> AddRoleToPlayerResult.PLAYER_NOT_FOUND;
                case ROLE_NOT_FOUND -> AddRoleToPlayerResult.ROLE_NOT_FOUND;
                case ALREADY_HAS_ROLE -> AddRoleToPlayerResult.ALREADY_HAS_ROLE;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }

    @Override
    public @NotNull RemoveRoleFromPlayerResult removeRoleFromPlayer(@NotNull UUID playerId, @NotNull String roleId) {
        var request = PermissionProto.RemoveRoleFromPlayerRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setRoleId(roleId)
                .build();

        try {
            this.grpc.removeRoleFromPlayer(request);
            return RemoveRoleFromPlayerResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PermissionProto.RemoveRoleFromPlayerError.class);

            return switch (error.getErrorType()) {
                case PLAYER_NOT_FOUND -> RemoveRoleFromPlayerResult.PLAYER_NOT_FOUND;
                case DOES_NOT_HAVE_ROLE -> RemoveRoleFromPlayerResult.DOES_NOT_HAVE_ROLE;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }
}
