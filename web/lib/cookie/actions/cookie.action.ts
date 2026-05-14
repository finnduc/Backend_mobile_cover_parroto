"use server";

import { cookies } from "next/headers";
import { AUTH_TOKEN_COOKIE } from "@/lib/constants";

const authCookieOptions = {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite: "lax" as const,
    path: "/",
    maxAge: 60 * 60,
};

export async function setAuthCookie(token: string) {
    const cookieStore = await cookies();
    cookieStore.set(AUTH_TOKEN_COOKIE, token, authCookieOptions);
}

export async function deleteAuthCookie() {
    const cookieStore = await cookies();
    cookieStore.delete(AUTH_TOKEN_COOKIE);
}