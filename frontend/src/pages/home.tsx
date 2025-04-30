import clsx from 'clsx';
import { toast, Toaster } from 'sonner';
import { compactLink } from '@/api/link';
import { Link } from '@/assets/components/link';
import { Input } from '@/assets/components/input';
import { Layout } from '@/assets/components/layout';
import { useState, useEffect, Fragment } from 'react';
import { Heart, Lock, Expand, QrCode } from 'lucide-react';
import { Dialog, DialogPanel, DialogBackdrop } from '@headlessui/react';
import { Field, FieldGroup, Label, ErrorMessage } from '@/assets/components/field';
import { Combobox, ComboboxLabel, ComboboxOption } from '@/assets/components/combobox';

const expirationOptions = [
   { value: '120', label: '2 Minutes' },
   { value: '300', label: '5 Minutes' },
   { value: '900', label: '15 Minutes' },
   { value: '1800', label: '30 Minutes' },
   { value: '3600', label: '1 Hour' }
];

export const Home = () => {
   const [isOpen, setIsOpen] = useState(false);
   const [showQR, setShowQR] = useState(false);
   const [prompt, setPrompt] = useState(undefined);

   const [linkName, setLinkName] = useState('');
   const [linkUUID, setLinkUUID] = useState('');

   const [formData, setFormData] = useState({ target: '', expire_in: '300' });

   const handleClose = () => {
      setIsOpen(false);
      setShowQR(false);
   };

   const handleClick = (event: any) => {
      if (event.altKey) {
         setShowQR(!showQR);
      } else {
         navigator.clipboard.writeText(`https://routekey.me/${linkName}`);
         toast('Copied to clipboard');
      }
   };

   const handleChange = (event: any) => {
      if (event.value) {
         setFormData({ ...formData, expire_in: event.value });
      } else {
         setFormData({ ...formData, [event.target.name]: event.target.value });
      }
   };

   const handleSubmit = (event: any) => {
      event.preventDefault();
      if (/^(ftp|http|https):\/\/[^ "]+$/.test(formData.target)) {
         compactLink(formData.target, formData.expire_in)
            .then((data: any) => {
               setLinkName(data.link);
               setLinkUUID(data.id);

               setPrompt('');
               setIsOpen(true);
            })
            .catch(err => {
               compactLink(formData.target, formData.expire_in)
                  .then((data: any) => {
                     setLinkName(data.link);
                     setLinkUUID(data.id);

                     setPrompt('');
                     setIsOpen(true);
                  })
                  .catch(err => {
                     console.error(err);
                     setPrompt('');
                  });
               setPrompt('');
            });
      } else {
         setPrompt('Please enter a valid url.');
      }
   };

   const Header = () => (
      <section className="text-center">
         <h2 className="select-none text-stone-500 text-lg font-medium mb-1">RouteKey: A url shortener built for the school on Key Route</h2>
         <hr className="w-[30rem] mx-auto border-stone-300" />
      </section>
   );

   const Footer = () => (
      <section className="text-center">
         <a
            href="https://github.com/theMackabu"
            target="_blank"
            rel="noopener noreferrer"
            className="group select-none inline-flex text-sm font-medium text-stone-600 hover:text-stone-700 transition">
            Made with <Heart className="mt-0.5 mx-1 stroke-transparent fill-red-400 h-4 w-4 group-hover:fill-red-500 transition" /> in Albany,
            California
         </a>
      </section>
   );

   return (
      <Layout header={<Header />} footer={<Footer />}>
         <Toaster />

         <Link
            to="/_/admin"
            className="absolute bottom-2 right-2 px-3 py-1.5 bg-black shadow-lg rounded-full text-white text-xs font-semibold tracking-tight opacity-85 hover:opacity-100 transition inset-ring-1 inset-ring-white/10 select-none">
            Have an ausdk12 staff account?
         </Link>

         <div className="absolute bottom-1 left-2 text-[8px] text-stone-400 font-serif select-none">{linkUUID}</div>

         <span className="text-9xl font-serif tracking-tight font-medium text-[#3D3D3A] flex items-center justify-center select-none">
            RouteKey <span className="mt-8 ml-1.5 text-6xl text-[#7957D9]">v2</span>
         </span>

         <p className="mt-14 text-sm text-center text-stone-600 select-none">
            Enter url to generate a routekey (automatically expires in 5 minutes by default)
         </p>

         <div className="mt-5 flex justify-center items-center">
            <form onSubmit={handleSubmit}>
               <FieldGroup className="w-[560px]">
                  <Field>
                     <Label> Full URL</Label>
                     <Input
                        type="text"
                        id="target"
                        name="target"
                        required
                        value={formData.target}
                        onChange={handleChange}
                        invalid={!!prompt}
                        autoComplete="off"
                        placeholder="https://blog.interviewing.io/can-fake-names-create-bias/"
                     />

                     {prompt && <ErrorMessage>{prompt}</ErrorMessage>}
                  </Field>

                  <Field>
                     <Label>Expires in</Label>
                     <Combobox
                        name="expire_in"
                        options={expirationOptions}
                        onChange={handleChange}
                        displayValue={option => option?.label}
                        defaultValue={expirationOptions[1]}>
                        {option => (
                           <ComboboxOption value={option}>
                              <ComboboxLabel>{option.label}</ComboboxLabel>
                           </ComboboxOption>
                        )}
                     </Combobox>
                  </Field>
               </FieldGroup>

               <div className="mt-5 flex flex-col justify-center">
                  <button
                     type="submit"
                     className="select-none w-full relative z-10 inline-flex min-h-[36px] cursor-pointer items-center justify-center border-0 bg-transparent px-3 pb-[0.3rem] text-base text-[#0d0b30] before:absolute before:inset-0 before:-z-10 before:block before:rounded-md before:border before:border-zinc-950/20 before:bg-white before:shadow-[0_1px_3px_0_rgba(20,20,96,0.1),inset_0_-5px_0_0_#ebebf6] before:content-[''] hover:before:border-[#8566fe] hover:before:bg-[#e0d9ff] hover:before:shadow-[0_2px_3px_0_rgba(20,20,96,0.1),inset_0_-5px_0_0_#c2b3ff] focus:outline-none focus-visible:before:outline focus-visible:before:outline-2 focus-visible:before:outline-[#8566fe] active:border-t-4 active:border-transparent active:py-1 active:before:shadow-none font-semibold">
                     Generate!
                  </button>
               </div>

               <Dialog open={isOpen} onClose={handleClose} className="relative z-50">
                  <DialogBackdrop
                     transition
                     className="fixed inset-0 flex w-screen justify-center overflow-y-auto bg-zinc-950/25 px-2 py-2 transition duration-100 focus:outline-0 data-closed:opacity-0 data-enter:ease-out data-leave:ease-in sm:px-6 sm:py-8 lg:px-8 lg:py-16 dark:bg-zinc-950/50 backdrop-blur-sm"
                  />

                  <div className="fixed inset-0 flex w-screen items-center justify-center">
                     <DialogPanel
                        transition
                        className="flex flex-col w-120 justify-between bg-stone-800 rounded-3xl p-8 mt-4 duration-300 ease-out data-closed:transform-[scale(95%)] data-closed:opacity-0">
                        <p className="text-lg sm:text-xl md:text-2xl text-white lg:w-1/2 select-none">Your generated RouteKey is</p>
                        <div className="flex flex-col gap-12 mt-20">
                           <div onClick={handleClick} className="flex flex-col cursor-pointer">
                              {showQR ? (
                                 <Fragment>
                                    <img
                                       className="-mt-9.5 rounded-lg size-40 p-2 bg-white shadow select-none pointer-events-none"
                                       src={`/api/v2/links/${linkUUID}/qrcode`}
                                    />
                                    <p className="text-base text-stone-400 mt-4 text-balance select-none">
                                       {'Go to '}
                                       <a className="text-indigo-300 select-text" href={`https://routekey.me/${linkName}`}>
                                          routekey.me/{linkName}
                                       </a>
                                       {' to use'}
                                    </p>
                                 </Fragment>
                              ) : (
                                 <Fragment>
                                    <p className="text-[clamp(1rem,5.4vw,12rem)] text-white font-semibold">{linkName}</p>
                                    <p className="text-base text-stone-400 mt-4 text-balance select-none">
                                       {'Go to '}
                                       <a className="text-indigo-300 select-text" href={`https://routekey.me/${linkName}`}>
                                          routekey.me/{linkName}
                                       </a>
                                       {' to use'}
                                    </p>
                                 </Fragment>
                              )}
                           </div>

                           <button
                              onClick={handleClose}
                              className="select-none flex transition rounded-full items-center duration-300 focus:outline-2 focus:ring-none ring-offset-zinc-950 focus:outline-offset-2 group relative text-black bg-white hover:text-accent-600 focus:outline-accent-600 h-14 px-6 py-3 text-lg font-medium pr-14">
                              <span className="z-10 pr-2">Again</span>
                              <div className="absolute right-1 inline-flex h-12 w-12 items-center justify-end rounded-full bg-stone-300 transition-[width] group-hover:w-[calc(100%-8px)]">
                                 <div className="relative h-12 w-12 flex items-center justify-center">
                                    <svg className="size-5" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                                       <path
                                          d="M8.14645 3.14645C8.34171 2.95118 8.65829 2.95118 8.85355 3.14645L12.8536 7.14645C13.0488 7.34171 13.0488 7.65829 12.8536 7.85355L8.85355 11.8536C8.65829 12.0488 8.34171 12.0488 8.14645 11.8536C7.95118 11.6583 7.95118 11.3417 8.14645 11.1464L11.2929 8H2.5C2.22386 8 2 7.77614 2 7.5C2 7.22386 2.22386 7 2.5 7H11.2929L8.14645 3.85355C7.95118 3.65829 7.95118 3.34171 8.14645 3.14645Z"
                                          fill="currentColor"
                                          fill-rule="evenodd"
                                          clip-rule="evenodd"></path>
                                    </svg>
                                 </div>
                              </div>
                           </button>
                        </div>
                     </DialogPanel>
                  </div>
               </Dialog>
            </form>
         </div>
      </Layout>
   );
};
