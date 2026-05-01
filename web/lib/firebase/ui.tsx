"use client";

import { ui } from "./client-app";
import { FirebaseUIProvider } from "@firebase-oss/ui-react";

export function FirebaseUIProviderHoc({ children }: { children: React.ReactNode }) {
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