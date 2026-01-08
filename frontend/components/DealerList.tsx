import { Dealer } from "../types";

interface DealerListProps {
  dealers: Dealer[];
}

export default function DealerList({ dealers }: DealerListProps) {
  return (
    <div className="grid gap-4">
      {dealers.length > 0 ? (
        dealers.map((dealer) => (
          <div 
            key={dealer.id} 
            className="flex items-center justify-between p-4 border rounded-lg border-zinc-200 dark:border-zinc-800 hover:bg-zinc-50 dark:hover:bg-zinc-800 transition-colors"
            data-testid="dealer-item"
          >
            <div>
              <h2 className="font-semibold text-zinc-900 dark:text-zinc-50">{dealer.name}</h2>
              <p className="text-sm text-zinc-600 dark:text-zinc-300">{dealer.city}</p>
            </div>
            <span className="text-xs font-semibold px-2.5 py-0.5 rounded bg-green-100 text-green-900 dark:bg-green-900 dark:text-green-100">
              Active
            </span>
          </div>
        ))
      ) : (
        <div className="p-8 text-center border-2 border-dashed rounded-lg border-zinc-300 dark:border-zinc-700">
          <p className="text-zinc-600 dark:text-zinc-400">Tidak ada data dealer ditemukan atau backend belum siap.</p>
        </div>
      )}
    </div>
  );
}
