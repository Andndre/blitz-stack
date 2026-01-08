import { Item } from "../types";
import ItemList from "../components/ItemList";

async function getItems(): Promise<Item[]> {
  const apiUrl = process.env.API_URL || "http://localhost:8080";
  console.log("Fetching from:", apiUrl);

  try {
    const res = await fetch(`${apiUrl}/api/items`, { next: { revalidate: 60 } });
    if (!res.ok) {
      throw new Error(`Failed to fetch: ${res.status}`);
    }
    return res.json();
  } catch (error) {
    console.error("Fetch error:", error);
    return [];
  }
}

export default async function Home() {
  const items = await getItems();

  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-zinc-50 p-8 font-sans dark:bg-black">
      <main className="flex w-full max-w-2xl flex-col gap-8 bg-white p-8 shadow-sm dark:bg-zinc-900 rounded-xl">
        <h1 className="text-3xl font-bold tracking-tight text-zinc-900 dark:text-zinc-50">
          Blitz Stack
        </h1>
        
        <p className="text-zinc-700 dark:text-zinc-300">
          Welcome to your new project. Data below is fetched from the Go backend:
        </p>

        <ItemList items={items} />

        <div className="pt-4 border-t border-zinc-200 dark:border-zinc-800">
          <p className="text-xs text-zinc-500 dark:text-zinc-400">
            Powered by Next.js + Bun + Go + PostgreSQL
          </p>
        </div>
      </main>
    </div>
  );
}
