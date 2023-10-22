import { RepoListing } from "@/components/molecules";
import { AppContext, type AppContextType } from "@/context/AppContext";
import { useContext } from "react";

export const RepoList = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className="">
      {ctx.repos?.map((item, i) => {
        return <RepoListing {...item} key={i} />;
      })}
    </div>
  );
};
