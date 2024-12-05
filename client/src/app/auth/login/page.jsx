import React from 'react'

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";


const page = () => {
  return (
    <section className="pb-32">
    <div className="container">
      <div className="flex flex-col gap-4">
        <div className="relative flex flex-col items-center overflow-hidden pb-2 pt-32">
          <img
            src="/site-logo.png"
            alt="logo"
            className="mb-4 h-10 w-auto"
          />
          <p className="text-2xl font-bold">Login</p>
        </div>
        <div className="z-10 mx-auto w-full max-w-sm rounded-md bg-background px-6 py-12 shadow">
          <div>
            <div className="grid gap-4">
              <div className="grid w-full max-w-sm items-center gap-1.5">
                <Label htmlFor="email">Email</Label>
                <Input type="email" placeholder="Enter your email" required />
              </div>
              <div>
                <div className="grid w-full max-w-sm items-center gap-1.5">
                  <Label htmlFor="password">Password</Label>
                  <Input
                    id="password"
                    type="password"
                    placeholder="Enter your password"
                    required
                  />
                </div>
              </div>
              <Button type="submit" className="mt-2 w-full">
                Login
              </Button>
            </div>
          </div>
        </div>
        <div className="mx-auto mt-3 flex justify-center gap-1 text-sm text-muted-foreground">
          <p>Don&apos;t have an account?</p>
          <a href="#" className="font-medium text-primary">
            Sign up
          </a>
        </div>
      </div>
    </div>
  </section>
  )
}

export default page



