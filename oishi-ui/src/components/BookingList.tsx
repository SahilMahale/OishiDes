import { Bookings } from '@/API/api';
import Table from './Table';

const BookingList = ({ data }: { data: Bookings[] }) => {
  const columns = ['User', 'BookingID'];
  return (
    <div>
      <Table data={data} columns={columns} />
    </div>
  );
};

export default BookingList;
