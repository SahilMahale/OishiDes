export type TableInfo = {
  available: boolean,
  seats: number,
  nonVeg: boolean,
  ac: boolean,
  id: number
}
export type CardInfo = {
  totalTables: number,
  tables: Array<TableInfo>,
}
export const cardInfo: CardInfo = {
  totalTables: 10,
  tables: [
    {
      available: true,
      seats: 4,
      nonVeg: true,
      ac: true,
      id: 1
    },
    {
      available: false,
      seats: 4,
      nonVeg: false,
      ac: true,
      id: 2
    },
    {
      available: true,
      seats: 2,
      nonVeg: true,
      ac: true,
      id: 3
    },
    {
      available: true,
      seats: 4,
      nonVeg: false,
      ac: false,
      id: 4
    },
    {
      available: false,
      seats: 2,
      nonVeg: true,
      ac: true,
      id: 5
    },
    {
      available: true,
      seats: 4,
      nonVeg: true,
      ac: true,
      id: 6
    },
    {
      available: true,
      seats: 4,
      nonVeg: true,
      ac: true,
      id: 7
    },
    {
      available: true,
      seats: 4,
      nonVeg: true,
      ac: true,
      id: 8
    },
    {
      available: true,
      seats: 4,
      nonVeg: true,
      ac: true,
      id: 9
    },
    {
      available: true,
      seats: 4,
      nonVeg: true,
      ac: true,
      id: 10
    },
  ],

}

