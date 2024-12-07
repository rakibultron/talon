import React from "react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import RegisterForm from "./components/RegisterForm";

const RegisterPage = () => {
  return (
    <section className="h-screen flex items-center justify-center bg-gray-50">
      <div className="w-full max-w-md p-6 bg-white rounded-md shadow-md">
        {/* Logo Section */}
        <div className="mb-6 text-center">
          <img
            src="/site-logo.png" // Replace with your logo URL
            alt="Logo"
            className="mx-auto h-16 w-16"
          />
          <h1 className="mt-4 text-2xl font-bold">Create an Account</h1>
          <p className="text-sm text-gray-500">Sign up to get started</p>
        </div>

        <RegisterForm />
      </div>
    </section>
  );
};

export default RegisterPage;
