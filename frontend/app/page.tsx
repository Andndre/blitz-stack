import { Dealer } from "../types";
import DealerList from "../components/DealerList";

async function getDealers(): Promise<Dealer[]> {
  const apiUrl = process.env.API_URL || "http://localhost:8080";
  console.log("Fetching from:", apiUrl);

  try {
    const res = await fetch(`${apiUrl}/api/dealers`, { next: { revalidate: 60 } });
    if (!res.ok) {
      throw new Error(`Failed to fetch data: ${res.status}`);
    }
    const data = await res.json();
    console.log("Fetch success, count:", data.length);
    return data;
  } catch (error) {
    console.error("Fetch error:", error);
    return [];
  }
}

export default async function Home() {
  const dealers = await getDealers();

  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-zinc-50 p-8 font-sans dark:bg-black">
      <main className="flex w-full max-w-2xl flex-col gap-8 bg-white p-8 shadow-sm dark:bg-zinc-900 rounded-xl">
        <h1 className="text-3xl font-bold tracking-tight text-zinc-900 dark:text-zinc-50">
          Dealer Heronusa
        </h1>
        
        <p className="text-zinc-700 dark:text-zinc-300">
          Data berikut diambil langsung dari Backend Go + Database PostgreSQL:
        </p>

        <DealerList dealers={dealers} />

        <div className="pt-4 border-t border-zinc-200 dark:border-zinc-800">
          <p className="text-xs text-zinc-500 dark:text-zinc-400">
            Stack: Next.js (Bun) + Go + PostgreSQL + Docker
          </p>
        </div>
      </main>
    </div>
  );
}