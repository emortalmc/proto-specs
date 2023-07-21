// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: game_tracker/models.proto

package gametracker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{0}
}

type BasicGamePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *BasicGamePlayer) Reset() {
	*x = BasicGamePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicGamePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicGamePlayer) ProtoMessage() {}

func (x *BasicGamePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicGamePlayer.ProtoReflect.Descriptor instead.
func (*BasicGamePlayer) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{1}
}

func (x *BasicGamePlayer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BasicGamePlayer) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CommonGameFinishWinnerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winners []*BasicGamePlayer `protobuf:"bytes,1,rep,name=winners,proto3" json:"winners,omitempty"`
	// losers may be empty if the game doesn't define losers
	Losers []*BasicGamePlayer `protobuf:"bytes,2,rep,name=losers,proto3" json:"losers,omitempty"`
}

func (x *CommonGameFinishWinnerData) Reset() {
	*x = CommonGameFinishWinnerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonGameFinishWinnerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonGameFinishWinnerData) ProtoMessage() {}

func (x *CommonGameFinishWinnerData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonGameFinishWinnerData.ProtoReflect.Descriptor instead.
func (*CommonGameFinishWinnerData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{2}
}

func (x *CommonGameFinishWinnerData) GetWinners() []*BasicGamePlayer {
	if x != nil {
		return x.Winners
	}
	return nil
}

func (x *CommonGameFinishWinnerData) GetLosers() []*BasicGamePlayer {
	if x != nil {
		return x.Losers
	}
	return nil
}

type BaseTowerDefenceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BluePlayers []*BasicGamePlayer `protobuf:"bytes,1,rep,name=blue_players,json=bluePlayers,proto3" json:"blue_players,omitempty"`
	RedPlayers  []*BasicGamePlayer `protobuf:"bytes,2,rep,name=red_players,json=redPlayers,proto3" json:"red_players,omitempty"`
}

func (x *BaseTowerDefenceData) Reset() {
	*x = BaseTowerDefenceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseTowerDefenceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseTowerDefenceData) ProtoMessage() {}

func (x *BaseTowerDefenceData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseTowerDefenceData.ProtoReflect.Descriptor instead.
func (*BaseTowerDefenceData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{3}
}

func (x *BaseTowerDefenceData) GetBluePlayers() []*BasicGamePlayer {
	if x != nil {
		return x.BluePlayers
	}
	return nil
}

func (x *BaseTowerDefenceData) GetRedPlayers() []*BasicGamePlayer {
	if x != nil {
		return x.RedPlayers
	}
	return nil
}

type TowerDefenceHealthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxHealth  int32 `protobuf:"varint,1,opt,name=max_health,json=maxHealth,proto3" json:"max_health,omitempty"`
	BlueHealth int32 `protobuf:"varint,2,opt,name=blue_health,json=blueHealth,proto3" json:"blue_health,omitempty"`
	RedHealth  int32 `protobuf:"varint,3,opt,name=red_health,json=redHealth,proto3" json:"red_health,omitempty"`
}

func (x *TowerDefenceHealthData) Reset() {
	*x = TowerDefenceHealthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerDefenceHealthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerDefenceHealthData) ProtoMessage() {}

func (x *TowerDefenceHealthData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerDefenceHealthData.ProtoReflect.Descriptor instead.
func (*TowerDefenceHealthData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{4}
}

func (x *TowerDefenceHealthData) GetMaxHealth() int32 {
	if x != nil {
		return x.MaxHealth
	}
	return 0
}

func (x *TowerDefenceHealthData) GetBlueHealth() int32 {
	if x != nil {
		return x.BlueHealth
	}
	return 0
}

func (x *TowerDefenceHealthData) GetRedHealth() int32 {
	if x != nil {
		return x.RedHealth
	}
	return 0
}

type TowerDefenceStartData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseData   *BaseTowerDefenceData   `protobuf:"bytes,1,opt,name=base_data,json=baseData,proto3" json:"base_data,omitempty"`
	HealthData *TowerDefenceHealthData `protobuf:"bytes,2,opt,name=health_data,json=healthData,proto3" json:"health_data,omitempty"`
}

func (x *TowerDefenceStartData) Reset() {
	*x = TowerDefenceStartData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerDefenceStartData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerDefenceStartData) ProtoMessage() {}

func (x *TowerDefenceStartData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerDefenceStartData.ProtoReflect.Descriptor instead.
func (*TowerDefenceStartData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{5}
}

func (x *TowerDefenceStartData) GetBaseData() *BaseTowerDefenceData {
	if x != nil {
		return x.BaseData
	}
	return nil
}

func (x *TowerDefenceStartData) GetHealthData() *TowerDefenceHealthData {
	if x != nil {
		return x.HealthData
	}
	return nil
}

type BlockSumoUpdateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scoreboard *BlockSumoScoreboard `protobuf:"bytes,1,opt,name=scoreboard,proto3" json:"scoreboard,omitempty"`
}

func (x *BlockSumoUpdateData) Reset() {
	*x = BlockSumoUpdateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSumoUpdateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSumoUpdateData) ProtoMessage() {}

func (x *BlockSumoUpdateData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSumoUpdateData.ProtoReflect.Descriptor instead.
func (*BlockSumoUpdateData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{6}
}

func (x *BlockSumoUpdateData) GetScoreboard() *BlockSumoScoreboard {
	if x != nil {
		return x.Scoreboard
	}
	return nil
}

type BlockSumoScoreboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlockSumoScoreboard) Reset() {
	*x = BlockSumoScoreboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSumoScoreboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSumoScoreboard) ProtoMessage() {}

func (x *BlockSumoScoreboard) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSumoScoreboard.ProtoReflect.Descriptor instead.
func (*BlockSumoScoreboard) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{7}
}

type MinesweeperStartData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalMines int64 `protobuf:"varint,1,opt,name=total_mines,json=totalMines,proto3" json:"total_mines,omitempty"` // TODO maybe we have a difficulty here?
}

func (x *MinesweeperStartData) Reset() {
	*x = MinesweeperStartData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinesweeperStartData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinesweeperStartData) ProtoMessage() {}

func (x *MinesweeperStartData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinesweeperStartData.ProtoReflect.Descriptor instead.
func (*MinesweeperStartData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{8}
}

func (x *MinesweeperStartData) GetTotalMines() int64 {
	if x != nil {
		return x.TotalMines
	}
	return 0
}

// MinesweeperUpdateData
// * You can calculate game progress by doing ((uncovered_mines) / total_mines) * 100. Note that total_mines is not in this message.
type MinesweeperUpdateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemainingMines int32 `protobuf:"varint,1,opt,name=remaining_mines,json=remainingMines,proto3" json:"remaining_mines,omitempty"`
	PlacedFlags    int32 `protobuf:"varint,2,opt,name=placed_flags,json=placedFlags,proto3" json:"placed_flags,omitempty"`
}

func (x *MinesweeperUpdateData) Reset() {
	*x = MinesweeperUpdateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinesweeperUpdateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinesweeperUpdateData) ProtoMessage() {}

func (x *MinesweeperUpdateData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinesweeperUpdateData.ProtoReflect.Descriptor instead.
func (*MinesweeperUpdateData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{9}
}

func (x *MinesweeperUpdateData) GetRemainingMines() int32 {
	if x != nil {
		return x.RemainingMines
	}
	return 0
}

func (x *MinesweeperUpdateData) GetPlacedFlags() int32 {
	if x != nil {
		return x.PlacedFlags
	}
	return 0
}

// MinesweeperFinishData
type MinesweeperFinishData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// winner_data if the game is won, all players are winners. If the game is lost, all players are losers.
	WinnerData *CommonGameFinishWinnerData `protobuf:"bytes,1,opt,name=winner_data,json=winnerData,proto3" json:"winner_data,omitempty"`
}

func (x *MinesweeperFinishData) Reset() {
	*x = MinesweeperFinishData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinesweeperFinishData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinesweeperFinishData) ProtoMessage() {}

func (x *MinesweeperFinishData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinesweeperFinishData.ProtoReflect.Descriptor instead.
func (*MinesweeperFinishData) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{10}
}

func (x *MinesweeperFinishData) GetWinnerData() *CommonGameFinishWinnerData {
	if x != nil {
		return x.WinnerData
	}
	return nil
}

type BlockSumoScoreboard_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player         *BasicGamePlayer `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	RemainingLives int32            `protobuf:"varint,2,opt,name=remaining_lives,json=remainingLives,proto3" json:"remaining_lives,omitempty"`
	Kills          int32            `protobuf:"varint,3,opt,name=kills,proto3" json:"kills,omitempty"`
	FinalKills     int32            `protobuf:"varint,4,opt,name=final_kills,json=finalKills,proto3" json:"final_kills,omitempty"`
}

func (x *BlockSumoScoreboard_Entry) Reset() {
	*x = BlockSumoScoreboard_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_models_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSumoScoreboard_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSumoScoreboard_Entry) ProtoMessage() {}

func (x *BlockSumoScoreboard_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_models_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSumoScoreboard_Entry.ProtoReflect.Descriptor instead.
func (*BlockSumoScoreboard_Entry) Descriptor() ([]byte, []int) {
	return file_game_tracker_models_proto_rawDescGZIP(), []int{7, 0}
}

func (x *BlockSumoScoreboard_Entry) GetPlayer() *BasicGamePlayer {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *BlockSumoScoreboard_Entry) GetRemainingLives() int32 {
	if x != nil {
		return x.RemainingLives
	}
	return 0
}

func (x *BlockSumoScoreboard_Entry) GetKills() int32 {
	if x != nil {
		return x.Kills
	}
	return 0
}

func (x *BlockSumoScoreboard_Entry) GetFinalKills() int32 {
	if x != nil {
		return x.FinalKills
	}
	return 0
}

var File_game_tracker_models_proto protoreflect.FileDescriptor

var file_game_tracker_models_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x06, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x22,
	0x3d, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa8,
	0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a,
	0x07, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x06, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x14, 0x42, 0x61,
	0x73, 0x65, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x4e, 0x0a, 0x0c, 0x62, 0x6c, 0x75, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x22, 0x77, 0x0a, 0x16, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x78, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x75,
	0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x62, 0x6c, 0x75, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x64, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x72, 0x65, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x54, 0x6f,
	0x77, 0x65, 0x72, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x53, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x75, 0x6d, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4f,
	0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x75, 0x6d, 0x6f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22,
	0xc4, 0x01, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x75, 0x6d, 0x6f, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0xac, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x43, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x76, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x4b, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x37, 0x0a, 0x14, 0x4d, 0x69, 0x6e, 0x65, 0x73, 0x77,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x69, 0x6e, 0x65, 0x73, 0x22,
	0x63, 0x0a, 0x15, 0x4d, 0x69, 0x6e, 0x65, 0x73, 0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x22, 0x70, 0x0a, 0x15, 0x4d, 0x69, 0x6e, 0x65, 0x73, 0x77, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x57, 0x0a,
	0x0b, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x77, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x60, 0x0a, 0x21, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_tracker_models_proto_rawDescOnce sync.Once
	file_game_tracker_models_proto_rawDescData = file_game_tracker_models_proto_rawDesc
)

func file_game_tracker_models_proto_rawDescGZIP() []byte {
	file_game_tracker_models_proto_rawDescOnce.Do(func() {
		file_game_tracker_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_tracker_models_proto_rawDescData)
	})
	return file_game_tracker_models_proto_rawDescData
}

var file_game_tracker_models_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_game_tracker_models_proto_goTypes = []interface{}{
	(*Game)(nil),                       // 0: emortal.model.game_tracker.Game
	(*BasicGamePlayer)(nil),            // 1: emortal.model.game_tracker.BasicGamePlayer
	(*CommonGameFinishWinnerData)(nil), // 2: emortal.model.game_tracker.CommonGameFinishWinnerData
	(*BaseTowerDefenceData)(nil),       // 3: emortal.model.game_tracker.BaseTowerDefenceData
	(*TowerDefenceHealthData)(nil),     // 4: emortal.model.game_tracker.TowerDefenceHealthData
	(*TowerDefenceStartData)(nil),      // 5: emortal.model.game_tracker.TowerDefenceStartData
	(*BlockSumoUpdateData)(nil),        // 6: emortal.model.game_tracker.BlockSumoUpdateData
	(*BlockSumoScoreboard)(nil),        // 7: emortal.model.game_tracker.BlockSumoScoreboard
	(*MinesweeperStartData)(nil),       // 8: emortal.model.game_tracker.MinesweeperStartData
	(*MinesweeperUpdateData)(nil),      // 9: emortal.model.game_tracker.MinesweeperUpdateData
	(*MinesweeperFinishData)(nil),      // 10: emortal.model.game_tracker.MinesweeperFinishData
	(*BlockSumoScoreboard_Entry)(nil),  // 11: emortal.model.game_tracker.BlockSumoScoreboard.Entry
}
var file_game_tracker_models_proto_depIdxs = []int32{
	1, // 0: emortal.model.game_tracker.CommonGameFinishWinnerData.winners:type_name -> emortal.model.game_tracker.BasicGamePlayer
	1, // 1: emortal.model.game_tracker.CommonGameFinishWinnerData.losers:type_name -> emortal.model.game_tracker.BasicGamePlayer
	1, // 2: emortal.model.game_tracker.BaseTowerDefenceData.blue_players:type_name -> emortal.model.game_tracker.BasicGamePlayer
	1, // 3: emortal.model.game_tracker.BaseTowerDefenceData.red_players:type_name -> emortal.model.game_tracker.BasicGamePlayer
	3, // 4: emortal.model.game_tracker.TowerDefenceStartData.base_data:type_name -> emortal.model.game_tracker.BaseTowerDefenceData
	4, // 5: emortal.model.game_tracker.TowerDefenceStartData.health_data:type_name -> emortal.model.game_tracker.TowerDefenceHealthData
	7, // 6: emortal.model.game_tracker.BlockSumoUpdateData.scoreboard:type_name -> emortal.model.game_tracker.BlockSumoScoreboard
	2, // 7: emortal.model.game_tracker.MinesweeperFinishData.winner_data:type_name -> emortal.model.game_tracker.CommonGameFinishWinnerData
	1, // 8: emortal.model.game_tracker.BlockSumoScoreboard.Entry.player:type_name -> emortal.model.game_tracker.BasicGamePlayer
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_game_tracker_models_proto_init() }
func file_game_tracker_models_proto_init() {
	if File_game_tracker_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_tracker_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicGamePlayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonGameFinishWinnerData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseTowerDefenceData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerDefenceHealthData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerDefenceStartData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSumoUpdateData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSumoScoreboard); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinesweeperStartData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinesweeperUpdateData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinesweeperFinishData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_tracker_models_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSumoScoreboard_Entry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_tracker_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_tracker_models_proto_goTypes,
		DependencyIndexes: file_game_tracker_models_proto_depIdxs,
		MessageInfos:      file_game_tracker_models_proto_msgTypes,
	}.Build()
	File_game_tracker_models_proto = out.File
	file_game_tracker_models_proto_rawDesc = nil
	file_game_tracker_models_proto_goTypes = nil
	file_game_tracker_models_proto_depIdxs = nil
}