import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { ShoppingCart, User, Menu, Heart } from "lucide-react";

const Header = () => {
  return (
    <header className="sticky top-0 z-50 border-b bg-white shadow-sm">
      <div className="container mx-auto flex items-center justify-between p-4">
        {/* Logo */}
        <div className="flex items-center gap-4">
          <Button variant="ghost" size="icon" className="lg:hidden text-black">
            <Menu />
          </Button>
          <a href="/" className="text-xl font-bold text-black">
            ShopLogo
          </a>
        </div>

        {/* Navigation */}
        <nav className="hidden lg:flex items-center space-x-6">
          <a href="/shop" className={cn("hover:text-gray-600", "text-black")}>
            Shop
          </a>
          <a href="/about" className={cn("hover:text-gray-600", "text-black")}>
            About Us
          </a>
          <a
            href="/contact"
            className={cn("hover:text-gray-600", "text-black")}
          >
            Contact
          </a>
          <a href="/blog" className={cn("hover:text-gray-600", "text-black")}>
            Blog
          </a>
        </nav>

        {/* Search Bar */}
        <div className="hidden md:flex w-1/3">
          <Input
            type="text"
            placeholder="Search products..."
            className="rounded-full"
          />
        </div>

        {/* Actions */}
        <div className="flex items-center space-x-4">
          <Button
            variant="ghost"
            size="icon"
            aria-label="Shopping Cart"
            className="relative text-black hover:text-gray-600"
          >
            <ShoppingCart />
            <span className="absolute top-0 right-0 inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white bg-blue-500 rounded-full">
              5
            </span>
          </Button>

          <Button
            variant="ghost"
            size="icon"
            aria-label="Wishlist"
            className="relative text-black hover:text-gray-600"
          >
            <Heart />
            <span className="absolute top-0 right-0 inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white bg-red-500 rounded-full">
              2
            </span>
          </Button>
          <Button
            variant="ghost"
            size="icon"
            aria-label="User Account"
            className="text-black hover:text-gray-600"
          >
            <User />
          </Button>
        </div>
      </div>
    </header>
  );
};

export default Header;
