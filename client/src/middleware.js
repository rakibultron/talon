import { NextResponse } from 'next/server';
import { jwtDecode } from "jwt-decode";
import { cookies } from 'next/headers';



const protectedPaths = ['/dashboard'];

export async function middleware(req) {

    const url = req.nextUrl.clone();

    // Get cookies
    const cookieStore = await cookies();
    const token = cookieStore.get('user_token')?.value;


    console.log({ token })
    // Check if the current URL matches any protected path
    if (protectedPaths.some(path => url.pathname.startsWith(path))) {
        // Redirect to login if no token
        if (!token) {
            console.log("No token found, redirecting to login.");
            return NextResponse.redirect(new URL('/auth/login', req.url));
        }

        // Decode and validate token
        try {
            const decoded = jwtDecode(token);

            // Check token expiration
            if (decoded.exp * 1000 < Date.now()) {
                console.log("Token expired, redirecting to login.");
                return NextResponse.redirect(new URL('/auth/login', req.url));
            }

            // If valid, allow the request to continue
            return NextResponse.next();
        } catch (error) {
            console.error("Invalid token, redirecting to login:", error.message);
            return NextResponse.redirect(new URL('/auth/login', req.url));
        }
    }

    // Allow requests to non-protected paths
    return NextResponse.next();
}


export const config = {
    matcher: ['/dashboard/:path*'],
};
