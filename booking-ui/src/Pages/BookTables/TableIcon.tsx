

const statusColors = new Map<string, string>([
  ["available", "fill- slate - 400"],
  ["selected", "fill-cyan-400"],
  ["unavailable", "fill - amber - 600"],
  ["booked", "fill-green-400"],
])


export default function TableIcon({ status = "available" }) {
  let fillColor = statusColors.get(status)
  if (fillColor) {
    fillColor = ""
  }
  return (
    <svg className={`${fillColor} stroke-black drop-shadow-xl md:hover:shadow-cyan-400 self-center`}
      xmlns="http://www.w3.org/2000/svg"
      xmlSpace="preserve"
      width={200}
      height={120}
      viewBox="-50 0 450 320 "
    ><path d="M332.483 71.393H173.195c.003-.093.014-.184.014-.277V56.427c0-5.523-4.477-10-10-10s-10 4.477-10 10v14.688c0 .094.011.184.014.277h-16.99c.003-.093.014-.184.014-.277 0-5.523-4.477-10-10-10h-64.65c-5.523 0-10 4.477-10 10 0 .094.011.184.014.277H10c-5.523 0-10 4.477-10 10v34.453c0 5.523 4.477 10 10 10h11.242v160.21c0 5.523 4.477 10 10 10s10-4.477 10-10v-160.21h260v160.21c0 5.523 4.477 10 10 10s10-4.477 10-10v-160.21h11.242c5.523 0 10-4.477 10-10V81.393c-.001-5.523-4.478-10-10.001-10zm-10 34.453H20V91.393h302.483v14.453z" /><path d="M145.459 159.498h-56.29c-5.523 0-10 4.477-10 10s4.477 10 10 10h5.465c.157 5.385 4.562 9.706 9.985 9.706h2.696v86.853h-2.696c-5.523 0-10 4.477-10 10s4.477 10 10 10h25.391c5.523 0 10-4.477 10-10s-4.477-10-10-10h-2.695v-86.853h2.695c5.424 0 9.828-4.32 9.985-9.706h5.465c5.523 0 10-4.477 10-10s-4.478-10-10.001-10zm107.855 0h-56.291c-5.523 0-10 4.477-10 10s4.477 10 10 10h5.465c.157 5.385 4.562 9.706 9.985 9.706h2.695v86.853h-2.695c-5.523 0-10 4.477-10 10s4.477 10 10 10h25.391c5.523 0 10-4.477 10-10s-4.477-10-10-10h-2.696v-86.853h2.696c5.424 0 9.828-4.32 9.985-9.706h5.465c5.523 0 10-4.477 10-10s-4.477-10-10-10z" />
    </svg>
  )
}
