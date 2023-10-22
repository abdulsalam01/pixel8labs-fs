import { AppContext, type AppContextType } from "@/context/AppContext";
import { type Follower } from "@/types";
import { useContext, useEffect } from "react";

export const useGetFollower = () => {
  const ctx = useContext(AppContext) as AppContextType;

  useEffect(() => {
    if (ctx.userData !== undefined) {
      const token = localStorage.getItem("token");
      if (token === null) {
        fetch(ctx.userData.followers_url)
          .then(async (res) => {
            return await res.json();
          })
          .then((data: Follower[]) => {
            ctx.setFollower(data);
          })
          .catch((err) => {
            console.log(err);
          });
      } else {
        fetch(ctx.userData.followers_url, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
          .then(async (res) => {
            return await res.json();
          })
          .then((data: Follower[]) => {
            ctx.setFollower(data);
          })
          .catch((err) => {
            console.log(err);
          });
      }
    }
  }, [ctx.userData]);
};
