package com.example.app

import android.app.Application
import com.clerk.api.Clerk

class EngflixApp : Application() {
    override fun onCreate() {
        super.onCreate()
        Clerk.initialize(
            context = this,
            publishableKey = BuildConfig.CLERK_PUBLISHABLE_KEY
        )
    }
}
