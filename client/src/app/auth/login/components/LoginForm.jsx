"use client";

import React from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useForm } from "react-hook-form";

const LoginForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmit = (data) => console.log(data);

  return (
    <div>
      {/* Login Form */}
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid gap-4">
          <div>
            <Input
              type="email"
              placeholder="Enter your email"
              required
              {...register("email")}
            />
          </div>
          <div>
            <Input
              type="password"
              placeholder="Enter your password"
              required
              {...register("password")}
            />
            <p className="mt-1 text-right text-sm text-primary cursor-pointer">
              Forgot password?
            </p>
          </div>
          <Button type="submit" className="mt-2 w-full">
            Log in
          </Button>
        </div>
      </form>
      {/* Footer Links */}
      <div className="mt-6 text-center text-sm text-gray-500">
        <p>
          Don't have an account?
          <a href="#" className="font-medium text-primary">
            Sign up
          </a>
        </p>
      </div>
    </div>
  );
};

export default LoginForm;
