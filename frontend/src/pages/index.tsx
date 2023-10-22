import { Section } from "@/components/molecules";
import { Footer, Header, RepoList, Sidebar } from "@/components/organisms";
import { useGetFollower, useGetRepo, useGetUser } from "@/hooks";

export default function Home() {
  useGetUser();
  useGetFollower();
  useGetRepo();

  return (
    <div className="min-h-screen flex flex-col justify-between">
      <section>
        <Header />
        <div className="flex justify-center mt-6">
          <div className="px-4 grid lg:grid-cols-4 gap-4 container">
            {/* Sidebar section */}
            <div className="md:px-4 col-span-4 lg:col-span-1">
              <Sidebar />
            </div>

            {/* Repositories section */}
            <div className="lg:p-6 col-span-4 lg:col-span-3 bg-white lg:border rounded-lg">
              <Section />
              <RepoList />
            </div>
          </div>
        </div>
      </section>

      <Footer />
    </div>
  );
}
