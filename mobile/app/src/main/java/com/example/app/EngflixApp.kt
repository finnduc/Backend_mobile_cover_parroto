package com.example.app

import android.app.Application
import android.util.Log
import com.clerk.api.Clerk

class EngflixApp : Application() {
    override fun onCreate() {
        super.onCreate()
        val publishableKey = BuildConfig.CLERK_PUBLISHABLE_KEY
        if (publishableKey.isBlank() || publishableKey == "null" || publishableKey.contains("REPLACE_ME")) {
            Log.e(TAG, "Missing CLERK_PUBLISHABLE_KEY. Auth flows are disabled until local.properties is configured.")
            return
        }

        try {
            Clerk.initialize(
                context = this,
                publishableKey = publishableKey
            )
        } catch (exception: IllegalArgumentException) {
            Log.e(TAG, "Invalid CLERK_PUBLISHABLE_KEY. Auth flows are disabled.", exception)
        }
    }

    private companion object {
        const val TAG = "EngflixApp"
    }
}
