import { AppProvider } from "@/context/AppContext";
import "@/styles/app.scss";
import type { AppProps } from "next/app";

import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <style jsx global>{`
        html {
          font-family: ${inter.style.fontFamily};
        }
      `}</style>
      <AppProvider>
        <Component {...pageProps} />
      </AppProvider>
    </>
  );
}
