import { type LoginResponse } from "@/types";
import { useSearchParams } from "next/navigation";
import { useEffect } from "react";

export const useGetToken = () => {
  const params = useSearchParams();

  useEffect(() => {
    if (params.get("code") != null) {
      const urlDest = `http://localhost:8000/api/callback?code=${params.get(
        "code",
      )}`;
      fetch(urlDest)
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
          console.log(err);
        });
    }
  }, [params]);
};
