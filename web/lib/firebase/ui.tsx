"use client";

import { useEffect } from "react";
import { onIdTokenChanged } from "firebase/auth";
import { FirebaseUIProvider } from "@firebase-oss/ui-react";
import { ui, auth } from "./client-app";
import { setAuthCookie, deleteAuthCookie } from "@/lib/cookie/actions/cookie.action";

export function FirebaseUIProviderHoc({ children }: { children: React.ReactNode }) {
  useEffect(() => {

    const unsubscribe = onIdTokenChanged(auth, async (user) => {
      if (user) {
        const token = await user.getIdToken();
        await setAuthCookie(token);
      } else {
        await deleteAuthCookie();
      }
    });

    return () => unsubscribe();
  }, []);

  return (
    <FirebaseUIProvider
      ui={ui}
      policies={{
        termsOfServiceUrl: "https://www.google.com",
        privacyPolicyUrl: "https://www.google.com",
      }}
    >
      {children}
    </FirebaseUIProvider>
  );
}