"use client";

import React from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useForm } from "react-hook-form";
import useAuth from "@/hooks/authHooks";
const RegisterForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const { userRegister } = useAuth();

  const onSubmit = async (data) => {
    const registerData = await userRegister("/auth/register", data);
    console.log({ registerData });
  };
  return (
    <div>
      {/* Registration Form */}
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid gap-4">
          {/* Name */}
          <Input
            type="text"
            placeholder="Full Name"
            required
            {...register("name")}
          />
          {/* Email */}
          <Input
            type="email"
            placeholder="Email Address"
            required
            {...register("email")}
          />
          {/* Password */}
          <Input
            type="password"
            placeholder="Password"
            required
            {...register("password")}
          />
          {/* Confirm Password */}
          <Input
            type="password"
            placeholder="Confirm Password"
            required
            {...register("confirmPassword")}
          />
          {/* Submit Button */}
          <Button type="submit" className="mt-2 w-full">
            Register
          </Button>
        </div>
      </form>
      {/* Footer Links */}
      <div className="mt-6 text-center text-sm text-gray-500">
        <p>
          Already have an account?
          <a href="#" className="font-medium text-primary">
            Log in
          </a>
        </p>
      </div>
    </div>
  );
};

export default RegisterForm;
