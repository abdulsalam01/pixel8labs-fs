import { Visitor } from "@/components/molecules";
import { AppContext, type AppContextType } from "@/context/AppContext";
import { useContext } from "react";

export const VisitorList = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className="mt-10">
      <h3 className="text-gray-800 font-bold text-sm">Latest Visitor</h3>
      <div className="flex items-center gap-4 mt-2">
        <Visitor item={ctx.followers?.[0]} />
        <Visitor item={ctx.followers?.[1]} />
        <Visitor item={ctx.followers?.[2]} />
      </div>
    </div>
  );
};
