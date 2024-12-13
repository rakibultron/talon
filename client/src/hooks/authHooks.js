import { useState, useEffect } from "react";
import axios from "@/lib/axios";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";

const useAuth = () => {
    const router = useRouter()
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);

    const userLogin = async (url, body, { headers = {}, params = {}, ...restOptions } = {}) => {
        try {
            setLoading(true);
            setError(null);
            const res = await axios.post(url, body, { headers, params, ...restOptions })
            console.log({ url })
            router.push("/")
            console.log({ res })
            return res.data;
        } catch (error) {
            // console.log({ error })
            setError(error);
            throw error;
        }
    }

    const userRegister = async (url, body, { headers = {}, params = {}, ...restOptions } = {}) => {
        try {
            const res = await axios.post(url, body, { headers, params, ...restOptions })
            router.push("/auth/login")
            return res.data;
        } catch (error) {
            toast(error.message)
            console.log({ error })
            setError(error);
            throw error;
        }
    }
    return { userLogin, userRegister, loading, error };
}

export default useAuth