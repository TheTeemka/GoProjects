import * as React from "react";
import { Button } from "@/components/ui/button";

export default function HomePage() {
  return (
    <div className="min-h-screen flex justify-center px-4">
      <div className="mt-45 w-full max-w-full rounded-lg text-center px-6 sm:px-12">
        <h1 className="text-6xl md:text-9xl font-bold mb-10 w-full">
          Welcome to <br />
          the App
        </h1>
        <p className="text-sm text-muted-foreground mb-4">
          Get started by signing in or creating a new account.
        </p>

        <div className="flex flex-col sm:flex-row gap-3 justify-center">
          <Button asChild className="w-full sm:w-auto">
            <a href="/login">Sign in</a>
          </Button>

          <Button asChild variant="outline" className="w-full sm:w-auto">
            <a href="/register">Create account</a>
          </Button>
        </div>
        <svg
          width="642"
          height="110"
          viewBox="0 0 642 110"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <svg
            width="642"
            height="110"
            viewBox="0 0 642 110"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          ></svg>
        </svg>
      </div>
    </div>
  );
}
