import { useState, useEffect } from "react";
import axios from "@/lib/axios";
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
            console.log({ error })
            setError(error);
            throw error;
        }
    }
    return { userLogin, loading, error };
}

export default useAuth