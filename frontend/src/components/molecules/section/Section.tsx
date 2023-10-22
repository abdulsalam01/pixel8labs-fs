import { BadgeMd, Heading } from "@/components/atoms";
import { AppContext, type AppContextType } from "@/context/AppContext";
import { useContext } from "react";

export const Section = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className="flex items-center gap-3 h-20">
      <Heading text="Repository" />
      <BadgeMd text={ctx.repos?.length.toString() ?? "0"} />
    </div>
  );
};
