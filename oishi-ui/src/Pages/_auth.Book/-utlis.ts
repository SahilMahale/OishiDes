
import { TABLESTATES } from "./-constants.ts"
export function checkAvailable(isAvailable = false) {
  return isAvailable ? TABLESTATES.AVAILABLE : TABLESTATES.UNAVAILABLE
}
