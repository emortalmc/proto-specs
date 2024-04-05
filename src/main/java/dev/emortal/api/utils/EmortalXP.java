package dev.emortal.api.utils;

public class EmortalXP {
    public static final double A = 0.15;
    public static final double B = 2.0;

    public static int toLevel(long xp) {
        return (int) (A * Math.pow(xp, 1/B));
    }

    public static long toXp(int level) {
        return (long) Math.ceil(Math.pow(level / A, B));
    }
}
