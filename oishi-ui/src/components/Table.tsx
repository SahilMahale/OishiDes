import { Bookings } from "@/API/api";

const Table = ({ data, columns }: { data: Bookings[], columns: Array<string> }) => {
  return (
    <div className="py-2 relative overflow-x-auto ">
      <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            {columns.map((columns, i) => (
              <th key={i} scope="col" className="px-6 py-3">
                {columns}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((tupl, i) => (
            <tr
              key={i}
              className="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
            >
              <th
                scope="row"
                className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
              >
                {tupl.user}
              </th>
              <td className="px-6 py-4">{tupl.bookingID}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};


export default Table;
