import React, { useReducer } from 'react';
//import { Button } from '@/components/ui/button';
import { TableCard } from './-TableCard';
import icon from './ButtonIcon.svg'
import { cardInfo, TableInfo } from './-mockAPI';
import { TABLESTATES, ACTIONS } from './-constants';
import { checkAvailable } from './-utlis';
import { useEffect, useMemo } from 'react';
import { useAuth } from '../../Context/AuthContext';
import { createFileRoute } from '@tanstack/react-router';

export type TableState = {
  tableInfo: TableInfo | undefined,
  tableStatus: string
}
export type TableGridState = {
  tablesMap: Map<number, TableState>
  selectedTables: Array<number>
}
export type TableStateActions = {
  type: string,
  tableInfo?: TableInfo,
  userName?: string
}

const initialState: TableGridState = {
  tablesMap: new Map<number, TableState>([
    [0, {
      tableInfo: undefined,
      tableStatus: TABLESTATES.AVAILABLE
    }],
  ]),
  selectedTables: [],
};
let reren = 0;


const BookTables = (bookInfo: TableGridState, selectedTablesArr: Array<number>) => {
  const selectedTablesInfo = new Map()
  selectedTablesArr.map((tableID) => {
    selectedTablesInfo.set(tableID, bookInfo.tablesMap.get(tableID))
  })
  const arr = []

  for (const tab of selectedTablesInfo.keys()) {
    arr.push(tab)
  }
  console.log("These tables are booked!", arr)
  return arr
}
// Refactor to match the new hashMap state
function tableReducer(tableState: TableGridState, action: TableStateActions): TableGridState {
  //console.log("DISPATCHING____________________________________________________")
  if (action.type === ACTIONS.SELECT) {
    const tableInfo = action.tableInfo;
    if (!tableInfo) {
      throw new Error("Error tableInfo is undefined")
    }
    const gridStateTableInfo = tableState.tablesMap.get(tableInfo.id)
    if (!tableInfo.available || !gridStateTableInfo) {
      alert("Table is already booked or is unavialable")
      return { ...tableState }
    }
    if (gridStateTableInfo.tableStatus === TABLESTATES.BOOKED) {
      alert("Table is already booked or is unavialable")
      return { ...tableState }
    }
    gridStateTableInfo.tableStatus = TABLESTATES.SELECTED;
    tableState.selectedTables.push(tableInfo.id)
    /* console.log("--------------------------------------------")
     console.log("AFTER SELECTING TABLE: ", tableInfo.id)
     console.log("tableSate:", tableState)
     console.log("--------------------------------------------")
     /*react doesn't rerender for same object with mutation,
      * it needs a new object like a babyi */
    return { ...tableState }
  } else if (action.type === ACTIONS.DESELECT) {
    const tableInfo = action.tableInfo;
    if (!tableInfo) {
      throw new Error("Error tableInfo is undefined")
    }
    const gridStateTableInfo = tableState.tablesMap.get(tableInfo.id)
    if (!gridStateTableInfo) {
      return { ...tableState }
    }
    gridStateTableInfo.tableStatus = TABLESTATES.AVAILABLE
    tableState.selectedTables = tableState.selectedTables.filter((elem) => {
      return elem !== tableInfo.id
    })
    return { ...tableState }
  } else if (action.type === ACTIONS.BOOK) {
    //change to use selectedTables array
    if (!action) {
      throw new Error("Error: Pass a valid userName")
    }
    if (action.userName == "") {
      throw new Error("Error: Pass a valid userName")
    }
    if (tableState.selectedTables?.length > 0) {
      const bookedTables = BookTables(tableState, tableState.selectedTables);
      console.log(typeof bookedTables, bookedTables)

      bookedTables.map((tId) => {
        // console.log("After fileter:", tableState.selectedTables)
        const tabelfromState = tableState.tablesMap.get(tId)
        if (!tabelfromState) {
          return
        }
        tabelfromState.tableStatus = TABLESTATES.BOOKED
      })

      tableState.selectedTables = []
      return { ...tableState }
    }
    return tableState
  }
  else if (action.type === ACTIONS.LOAD) {
    const tableInfo = action.tableInfo;
    //console.log("start loading")
    if (!tableInfo) {
      throw new Error("Error tableInfo is undefined")
    }
    tableState.tablesMap.set(tableInfo.id, {
      tableInfo: tableInfo,
      tableStatus: checkAvailable(tableInfo.available)
    })
    // Do not want reRenders here just bulk loading everything in  
    return tableState
  } else {
    alert("Select a valid table to be booked")
    return tableState
  }
}

export const Route = createFileRoute('/_auth/Book/Tables')({
  component: Book
})

function Book() {
  console.log("-----------------------------------------")
  reren++
  console.log("Re-Render Book page count:", reren)
  console.log("-----------------------------------------")
  const [tablesStates, dispatchTables] = useReducer<React.Reducer<TableGridState, TableStateActions>>(tableReducer, initialState);
  const { Context: appContext } = useAuth()
  const userName = appContext.claims.name
  // to avoid too many re-render
  useMemo(() => {
    cardInfo.tables.map((table) => {
      console.log("call load")
      try {
        dispatchTables({ type: ACTIONS.LOAD, tableInfo: table })
      } catch (e) {
        console.error(e)
      }
    })
  }, [])
  useEffect(() => {
    console.log("Cards selected", tablesStates.selectedTables)
  }, [tablesStates])
  return (
    <div className="bg-slate-900 min-h-screen">
      <div className="container mx-auto rounded-md pb-10 bg-slate-700 flex flex-col items-center ">
        <div className='drop-shadow-2xl'>
          < h2 className="font-sans text-slate-200 tracking-tight text-3xl text-center font-bold " >
            {`Book Tables for ${userName}`}
          </h2 >
        </div>
        <div className='flex flex-col'>
          <div className=' mt-50 lg:grid lg:grid-cols-5 gap-5 md:flex md:flex-col'>
            {cardInfo && (cardInfo.tables.map((table) => {
              const tState = tablesStates.tablesMap.get(table.id)
              if (!tState) {
                return
              }
              return (
                < TableCard key={table.id} tableInfo={table} state={tState} dispatcherFunc={dispatchTables} />
              )
            }))}
          </div>
          <div className='fixed lg:bottom-36 lg:right-44 sm:right-32 right-[1px] p-4 flex'>
            <button
              onClick={() => {
                dispatchTables({ type: ACTIONS.BOOK })
              }}
              className={` ${(tablesStates.selectedTables?.length > 0) ? "bg-cyan-400 text-cyan-900" : "bg-gray-500 text-gray-200"}   rounded-full   drop-shadow-2xl font-bold`}>
              <img src={icon} />
              Book
            </button>
          </div>
        </div>
      </div >
    </div >
  );
}

