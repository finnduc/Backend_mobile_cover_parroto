"use client";

import { autoAnonymousLogin, initializeUI } from "@firebase-oss/ui-core";
import { getApps, initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import { firebaseConfig } from "./config";

export const firebaseApp = getApps().length === 0 ? initializeApp(firebaseConfig) : getApps()[0];

export const auth = getAuth(firebaseApp);

export const ui = initializeUI({
  app: firebaseApp,
  // behaviors: [autoAnonymousLogin()],
});