'use client'

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { clientApiFetch } from '@/lib/client-api';
import { ROUTES } from '@/lib/routes';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';

const ROLES = [
  { value: 'student', label: 'Student', description: 'I want to learn and practice' },
  { value: 'teacher', label: 'Teacher', description: 'I want to create and manage lessons' },
];

export function OnboardingScreen() {
  const router = useRouter();
  const [selectedRole, setSelectedRole] = useState<string>('student');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleComplete = async () => {
    setLoading(true);
    setError(null);

    const res = await clientApiFetch('/complete-signup', {
      method: 'POST',
      body: JSON.stringify({ role: selectedRole }),
    });

    if (res.error) {
      setError(res.error.message);
      setLoading(false);
      return;
    }

    router.push(ROUTES.USER.HOME);
  };

  return (
    <div className="flex items-center justify-center min-h-screen">
      <Card className="max-w-md mx-auto">
        <CardHeader>
          <CardTitle>Welcome!</CardTitle>
          <CardDescription>Choose your role to get started</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            {ROLES.map((role) => (
              <button
                key={role.value}
                type="button"
                onClick={() => setSelectedRole(role.value)}
                className={`w-full text-left p-4 rounded-lg border-2 transition-colors ${
                  selectedRole === role.value
                    ? 'border-primary bg-primary/5'
                    : 'border-border hover:border-muted-foreground'
                }`}
              >
                <div className="font-medium">{role.label}</div>
                <div className="text-sm text-muted-foreground">{role.description}</div>
              </button>
            ))}
          </div>

          {error && (
            <p className="text-sm text-destructive">{error}</p>
          )}

          <Button
            className="w-full"
            onClick={handleComplete}
            disabled={loading}
          >
            {loading ? 'Setting up...' : 'Continue'}
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
