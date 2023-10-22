import { type LoginResponse } from "@/types";
import { useSearchParams } from "next/navigation";
import { useEffect } from "react";

export const useGetToken = () => {
  const params = useSearchParams();

  useEffect(() => {
    const url = `${process.env.NEXT_PUBLIC_BASE_URL}/callback?code=${params.get(
      "code",
    )}`;
    if (params.get("code") != null) {
      fetch(url)
        .then(async (res) => {
          return await res.json();
        })
        .then((data) => {
          const loginData = data.data as LoginResponse;
          localStorage.setItem("token", loginData.access_token);
        })
        .then(() => {
          window.location.href = "/";
        })
        .catch((err) => {
          throw err;
        });
    }
  }, [params]);
};
