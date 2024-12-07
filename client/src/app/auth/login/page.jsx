import React from "react";

import LoginForm from "./components/LoginForm";

const LoginPage = () => {
  return (
    <section className="h-screen flex items-center justify-center ">
      <div className="w-full max-w-md p-6 rounded-md shadow-md">
        {/* Logo Section */}
        <div className="mb-6 text-center">
          <img
            src="/site-logo.png" // Replace with your logo URL
            alt="Logo"
            className="mx-auto h-16 w-16"
          />
          <h1 className="mt-4 text-2xl font-bold">Welcome Back</h1>
          <p className="text-sm text-gray-500">Log in to your account</p>
        </div>

        <LoginForm />
      </div>
    </section>
  );
};

export default LoginPage;
