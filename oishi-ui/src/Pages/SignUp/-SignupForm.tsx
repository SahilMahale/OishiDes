import { useMutation } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-router';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { useForm } from 'react-hook-form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { userSignup } from '../../API/api';
import Loading from '../../components/Loading'

const signupFormSchema = z.object({
  user: z.string().min(4, { message: "UserName should be atleast 4 characters long" }).max(50),
  pass: z.string().min(8, { message: "password should be atleast 8 characters long" }),
  email: z.string().email({ message: "Please enter a valid email ID" })
})
export const SignupForm = ({ isAdmin = false }) => {
  const navigateTO = useNavigate();
  const { mutate, isPending, isPaused, isError, error } = useMutation({
    mutationFn: ({ user, email, pass }: { user: string, email: string, pass: string }) => {
      return userSignup(user, email, pass, isAdmin);
    },
    onSuccess: () => {
      navigateTO({ to: '/Login' });
    },
  });
  const signupForm = useForm<z.infer<typeof signupFormSchema>>({
    defaultValues: {
      user: "Monkey-menace",
      pass: '',
      email: 'monkey@saiyan.com',
    },
    resolver: zodResolver(signupFormSchema)
  })
  const handleSubmitSignup = (vals: z.infer<typeof signupFormSchema>) => {
    mutate({ user: vals.user, email: vals.email, pass: vals.pass });
  }
  return (
    <>
      {(isPending || isPaused) && (
        <div className="container mx-auto px-20 py-2 flex flex-col items-center ">
          <Loading />
        </div>
      )}
      {isError && (
        <div className='container mx-auto px-20 py-2 flex flex-col items-center '>
          <h3>{error.message}</h3>
        </div>
      )}
      {!(isPending || isPaused) && (
        <div className="container mx-auto px-20 py-2 flex flex-wrap items-center justify-between ">
          <Form
            {...signupForm}
          >
            <form className="pl-10" onSubmit={signupForm.handleSubmit(handleSubmitSignup)}>
              <div className=" mb-6 ">
                <FormField
                  control={signupForm.control}
                  name='user'
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel
                        className="block mb-2 text-lg font-medium text-slate-300"
                      >
                        User Name
                      </FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          className="bg-gray-50 border w-[300px] text-gray-900 text-lg rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 border-gray-600 placeholder-gray-400"
                        />
                      </FormControl>
                      <FormDescription className='text-slate-400'>
                        Enter the username which is atleast 4 characters long
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div className="mb-6">
                <FormField
                  control={signupForm.control}
                  name='email'
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel
                        className="block mb-2 text-lg font-medium text-slate-300"
                      >
                        Email address
                      </FormLabel>
                      <FormControl>
                        <input
                          {...field}
                          className="bg-gray-50 w-[300px] border  text-gray-900 text-lg rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 border-gray-600 placeholder-gray-400"
                        />
                      </FormControl>
                      <FormDescription className='text-slate-400'>
                        Enter a valid Email address
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              <div className="mb-6">
                <FormField
                  control={signupForm.control}
                  name='pass'
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel
                        className="block mb-2 text-lg font-medium text-slate-200"
                      >
                        Password
                      </FormLabel>
                      <FormControl>
                        <input
                          type='password'
                          {...field}
                          className=" w-[300px] border text-gray-900 text-lg rounded-lg  block p-2.5 border-gray-600 placeholder-gray-400 focus:ring-blue-500 focus:border-blue-500"
                        />
                      </FormControl>
                      <FormDescription className='text-slate-400'>
                        Enter a passWord 8 characters long
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>

              <Button
                type="submit"
                className="bg-transparent hover:bg-slate-500 text-slate-300 font-semibold hover:text-white py-2 px-4 border border-slate-700 hover:border-transparent rounded"
              >
                Sign Up
              </Button>
            </form>
          </Form>
          <div className="pr-20 py-2">
            {!isAdmin && (
              <svg
                aria-label="usersSVG"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 256 256"
                className="h-96 w-96 object-contain fill-slate-700 stroke-slate-800"
              >
                <path d="M57.2 80.8c8.7 0 16.1-3.1 22.3-9.2 6.1-6.1 9.2-13.6 9.2-22.3 0-8.7-3.1-16.1-9.2-22.3-6.1-6.1-13.6-9.2-22.3-9.2-8.7 0-16.1 3.1-22.3 9.2-6.1 6.1-9.2 13.6-9.2 22.3 0 8.7 3.1 16.1 9.2 22.3 6.2 6.1 13.6 9.2 22.3 9.2zM94.6 129.9c9.2 9.2 20.3 13.8 33.4 13.8 13 0 24.2-4.6 33.4-13.8 9.2-9.2 13.8-20.3 13.8-33.4 0-13-4.6-24.2-13.8-33.4-9.2-9.2-20.3-13.8-33.4-13.8-13 0-24.2 4.6-33.4 13.8-9.2 9.2-13.8 20.3-13.8 33.4 0 13.1 4.6 24.2 13.8 33.4zM198.8 80.8c8.7 0 16.1-3.1 22.3-9.2 6.1-6.1 9.2-13.6 9.2-22.3 0-8.7-3.1-16.1-9.2-22.3-6.1-6.1-13.6-9.2-22.3-9.2s-16.1 3.1-22.3 9.2-9.2 13.6-9.2 22.3c0 8.7 3.1 16.1 9.2 22.3 6.2 6.1 13.6 9.2 22.3 9.2zM230.8 80.8c-.5 0-2.3.9-5.3 2.6s-7.1 3.5-12 5.2c-4.9 1.8-9.8 2.6-14.6 2.6-5.5 0-10.9-.9-16.4-2.8.4 3 .6 5.7.6 8.1 0 11.4-3.3 21.9-10 31.5 13.3.4 24.1 5.7 32.6 15.7h16.5c6.7 0 12.4-1.7 17-5s6.9-8.2 6.9-14.6c-.1-28.8-5.2-43.3-15.3-43.3z" />
                <path d="M212.4 180.2c-.9-4.8-2-9.2-3.3-13.3s-3.1-8.1-5.3-12c-2.2-3.9-4.7-7.2-7.6-9.9-2.9-2.7-6.4-4.9-10.5-6.6-4.1-1.6-8.7-2.5-13.7-2.5-.8 0-2.6.9-5.3 2.6-2.7 1.8-5.7 3.7-9 5.9s-7.7 4.1-13.2 5.9-11 2.6-16.6 2.6c-5.6 0-11.1-.9-16.6-2.6s-9.9-3.7-13.1-5.9c-3.3-2.2-6.3-4.1-9-5.9-2.7-1.8-4.5-2.6-5.3-2.6-5 0-9.6.8-13.7 2.5-4.1 1.6-7.6 3.8-10.5 6.6-2.9 2.7-5.4 6.1-7.6 9.9-2.2 3.9-4 7.9-5.3 12-1.3 4.1-2.4 8.5-3.3 13.3-.9 4.8-1.4 9.2-1.7 13.4-.3 4.1-.4 8.4-.4 12.7 0 9.8 3 17.6 9 23.3 6 5.7 13.9 8.5 23.8 8.5h107.4c9.9 0 17.9-2.8 23.8-8.5 6-5.7 9-13.5 9-23.3 0-4.3-.1-8.6-.4-12.7-.2-4.2-.7-8.6-1.6-13.4zM82.9 128c-6.6-9.6-9.9-20.1-9.9-31.5 0-2.4.2-5.1.6-8.1-5.4 1.9-10.9 2.8-16.4 2.8-4.8 0-9.7-.9-14.6-2.6s-8.9-3.5-12-5.2c-3.1-1.7-4.9-2.6-5.3-2.6-10.2 0-15.3 14.5-15.3 43.4 0 6.4 2.3 11.2 6.9 14.6 4.6 3.3 10.2 5 17 5h16.5c8.4-10.1 19.2-15.4 32.5-15.8z" />
              </svg>
            )}
            {isAdmin && (
              <svg
                xmlns="http://www.w3.org/2000/svg"
                aria-label="adminSVG"
                xmlSpace="preserve"
                viewBox="0 0 512.001 512.001"
                className="h-96 w-96 object-contain fill-slate-700 stroke-slate-800"
              >
                <path d="M454.275 85.672c-68.879 0-116.732-22.029-144.749-40.51-30.433-20.075-44.786-40.46-44.913-40.641A10.45 10.45 0 0 0 256.02 0h-.027a10.475 10.475 0 0 0-8.589 4.499c-.141.204-14.495 20.589-44.927 40.664-28.016 18.481-75.87 40.51-144.748 40.51-5.771 0-10.449 4.678-10.449 10.449v158.065l.002.239c1.056 46.063 4.087 87.558 30.591 128.922 29.482 46.013 84.693 86.587 173.75 127.69a10.458 10.458 0 0 0 8.756 0c89.058-41.103 144.268-81.676 173.751-127.69 26.504-41.364 29.535-82.859 30.591-128.922l.002-.239V96.121c.001-5.771-4.677-10.449-10.448-10.449zm-10.449 168.396c-1.02 44.329-3.692 81.175-27.291 118.006-26.767 41.774-77.836 79.317-160.533 117.957-82.696-38.639-133.766-76.183-160.532-117.957-23.6-36.83-26.272-73.677-27.291-118.006V106.411c69.059-2.128 117.68-25.055 146.76-44.435 19.66-13.102 33.152-26.244 41.063-35.04 7.911 8.796 21.403 21.938 41.063 35.04 29.081 19.381 77.701 42.308 146.761 44.435v147.657z" />
                <path d="M208.563 99.101c-2.665-5.119-8.974-7.109-14.093-4.446-28.265 14.708-60.078 24.569-94.553 29.307-5.717.786-9.715 6.057-8.929 11.775.719 5.234 5.2 9.027 10.339 9.027.474 0 .954-.032 1.436-.098 36.855-5.066 70.955-15.654 101.354-31.472 5.119-2.664 7.109-8.974 4.446-14.093zM240.65 77.271c-3.424-4.647-9.965-5.637-14.61-2.214a214.754 214.754 0 0 1-3.758 2.708c-4.715 3.328-5.839 9.847-2.512 14.563a10.439 10.439 0 0 0 8.546 4.424c2.083 0 4.185-.621 6.016-1.913a255.667 255.667 0 0 0 4.103-2.958c4.647-3.423 5.638-9.964 2.215-14.61zM256.001 123.865c-74.902 0-135.84 60.938-135.84 135.841 0 74.902 60.938 135.839 135.84 135.839s135.84-60.937 135.84-135.839c0-74.903-60.938-135.841-135.84-135.841zm0 250.781c-36.545 0-69.16-17.145-90.228-43.813 19.88-18.358 45.799-30.142 73.897-33.436a10.45 10.45 0 0 0 2.207-20.251c-17.332-6.009-28.978-22.375-28.978-40.723 0-23.766 19.335-43.101 43.102-43.101 23.767 0 43.102 19.335 43.102 43.101 0 18.348-11.646 34.712-28.979 40.723a10.448 10.448 0 0 0 2.207 20.251c28.098 3.295 54.017 15.078 73.898 33.436-21.069 26.668-53.683 43.813-90.228 43.813zm44.94-92.661c11.926-11.776 19.06-28.061 19.06-45.561 0-35.289-28.71-63.999-64-63.999s-64 28.71-64 63.999c0 17.499 7.136 33.784 19.06 45.561-21.163 6.237-40.58 16.859-56.815 31.125-8.414-15.966-13.186-34.137-13.186-53.403 0-63.38 51.563-114.942 114.941-114.942s114.941 51.562 114.941 114.941c0 19.266-4.772 37.437-13.186 53.404-16.235-14.267-35.652-24.888-56.815-31.125z" />
              </svg>
            )}
          </div>
        </div>
      )}
    </>
  );
};
