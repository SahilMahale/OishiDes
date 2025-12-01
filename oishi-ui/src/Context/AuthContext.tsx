import { jwtDecode } from 'jwt-decode';
import React, {
  createContext,
  useContext,
  useMemo,
  useReducer,
  ReactNode,
  useCallback
} from 'react';
import { sleep } from './utils';

const enum ACTIONS {
  LOGIN,
  LOGOUT,
  RECOVER,
  RETURNSTATE
}
export type JWTclaims = {
  name: string,
  type: string
}

export type Context = {
  isLoggedIn: boolean,
  claims: JWTclaims,
  token: string,
  isRecovered: boolean,
};

export interface AppContext {
  Context: Context
  setToken?(jwtToken: string): void
  LogOut?(): Promise<boolean>
}
export type ActionType = {
  type: number,
  token?: string,
}

const initialContext: Context = {
  isLoggedIn: false,
  claims: { name: "dummy", type: "dummy" },
  token: '',
  isRecovered: false,
}

const reducer = (contextState: Context, action: ActionType): Context => {
  if (action.type === ACTIONS.LOGIN) {
    if (!action.token) {
      throw new Error("Error: the JWT was not Passed")
    }
    contextState.token = action.token;
    contextState.claims = jwtDecode(action.token);
    contextState.isLoggedIn = true;
    localStorage.setItem('authard', JSON.stringify(contextState));
    contextState.isRecovered = true;
    return { ...contextState };
  } else if (action.type === ACTIONS.RECOVER) {
    const localAuthKey = localStorage.getItem('authard')
    if (!localAuthKey) {
      //alert("No authKeyfound")
      contextState = initialContext;
      localStorage.setItem('authard', JSON.stringify(contextState));
      contextState.isRecovered = true;
      return { ...contextState }
    }
    contextState = JSON.parse(localAuthKey);
    console.log("-----CONTEXT RECOVERED: ", contextState)
    contextState.isRecovered = true;
    return { ...contextState };
  } else if (action.type === ACTIONS.LOGOUT) {
    contextState = initialContext;
    localStorage.setItem('authard', JSON.stringify(contextState));
    return { ...contextState };
  } else if (action.type === ACTIONS.RETURNSTATE) {
    return { ...contextState };
  } else {
    throw new Error('INVALID ACTION FOR REDUCER');
  }
};
const initialAppContext: AppContext = {
  Context: initialContext,
}

let contextVal: AppContext
export const AuthContext = createContext<AppContext>(initialAppContext);

const AuthProvider = ({ children }: { children: ReactNode }) => {

  const [appContext, dispatch] = useReducer<React.Reducer<Context, ActionType>>(reducer, initialContext);
  console.log("---------------AUTH RELOAD----------------")
  const SetToken = useCallback(async (jwtToken: string) => {
    try {
      dispatch({ type: ACTIONS.LOGIN, token: jwtToken });
    }
    catch (e) {
      console.error(e)
    }
    await sleep(500)
    //  window.location.replace('/');
  }, []);
  const LogOut = useCallback(async () => {
    try {
      dispatch({ type: ACTIONS.LOGOUT });
    }
    catch (e) {
      console.error(e)
    }
    return Promise.resolve(true)
  }, []);
  useMemo(() => {

    console.log("---------------AUTH RECOVER BEGIN----------------")
    dispatch({ type: ACTIONS.RECOVER });

  }, []);
  useMemo(
    () => {
      contextVal = {
        Context: appContext,
        setToken: SetToken,
        LogOut: LogOut,
      }
    }, [appContext, SetToken, LogOut]
  );
  return (
    <AuthContext.Provider value={contextVal}>{children}</AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
};

export default AuthProvider;
