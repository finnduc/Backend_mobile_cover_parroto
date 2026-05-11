'use client'

import Link from 'next/link'
import { Loader2, ArrowLeft } from 'lucide-react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

export function OnboardingScreen() {
    return (
        <div className="max-w-sm mx-auto text-center">
            <Card>
                <CardHeader>
                    <CardTitle>Setting up your account</CardTitle>
                    <CardDescription>Please wait while we configure your account.</CardDescription>
                </CardHeader>
                <CardContent>
                    <div className="flex flex-col items-center gap-4 py-8">
                        <Loader2 className="h-10 w-10 animate-spin text-primary" />
                        <p className="text-sm text-muted-foreground">
                            Đang thiết lập tài khoản, vui lòng chờ...
                        </p>
                        <Button variant="link" size="sm" asChild className="mt-2">
                            <Link href="/">
                                <ArrowLeft className="h-4 w-4 mr-1" />
                                Quay lại trang chủ
                            </Link>
                        </Button>
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}
