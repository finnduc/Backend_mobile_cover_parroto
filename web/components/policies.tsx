import Link from "next/link";

export function Policies() {
  return (
    <div className="text-text-muted text-center text-xs">
      By creating an account, you agree to our{" "}
      <Link
        href="https://www.google.com"
        target="_blank"
        rel="noopener noreferrer"
        className="hover:underline font-semibold"
      >
        Terms of Service
      </Link>{" "}
      and{" "}
      <Link
        href="https://www.google.com"
        target="_blank"
        rel="noopener noreferrer"
        className="hover:underline font-semibold"
      >
        Privacy Policy
      </Link>
      .
    </div>
  );
}
