import { useRouter } from "next/router";
import { useEffect } from "react";

export const useCheckLogin = () => {
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (token === null) {
      router.push("/octocat");
    }
  }, []);
};
