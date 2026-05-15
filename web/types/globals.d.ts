import { UserRole } from "@/lib/enums/user-role.enum"

export {}
declare global {
 interface CustomJwtSessionClaims {
 metadata: {
 role?: UserRole
 }
 }
}