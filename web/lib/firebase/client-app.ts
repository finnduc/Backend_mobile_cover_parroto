"use client";

import { initializeApp, getApps } from "firebase/app";
import { firebaseConfig } from "./config";
import { connectAuthEmulator, getAuth } from "firebase/auth";
import { autoAnonymousLogin, initializeUI } from "@firebase-oss/ui-core";

export const firebaseApp = getApps().length === 0 ? initializeApp(firebaseConfig) : getApps()[0];

export const auth = getAuth(firebaseApp);

export const ui = initializeUI({
  app: firebaseApp,
  behaviors: [autoAnonymousLogin()],
});

if (process.env.NODE_ENV === "development") {
  connectAuthEmulator(auth, "http://localhost:9099");
}