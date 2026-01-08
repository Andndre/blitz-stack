import { Item } from "../types";

interface ItemListProps {
  items: Item[];
}

export default function ItemList({ items }: ItemListProps) {
  return (
    <div className="grid gap-4">
      {items.length > 0 ? (
        items.map((item) => (
          <div 
            key={item.id} 
            className="flex items-center justify-between p-4 border rounded-lg border-zinc-200 dark:border-zinc-800 hover:bg-zinc-50 dark:hover:bg-zinc-800 transition-colors"
            data-testid="item-row"
          >
            <div>
              <h2 className="font-semibold text-zinc-900 dark:text-zinc-50">{item.title}</h2>
              <p className="text-sm text-zinc-600 dark:text-zinc-300">{item.description}</p>
            </div>
            <span className="text-xs font-semibold px-2.5 py-0.5 rounded bg-blue-100 text-blue-900 dark:bg-blue-900 dark:text-blue-100">
              Sample
            </span>
          </div>
        ))
      ) : (
        <div className="p-8 text-center border-2 border-dashed rounded-lg border-zinc-300 dark:border-zinc-700">
          <p className="text-zinc-600 dark:text-zinc-400">No items found. Run 'make seed' to populate data.</p>
        </div>
      )}
    </div>
  );
}
