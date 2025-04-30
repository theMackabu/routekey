import { Link } from '@/assets/components/link';
import { PageContentBlock } from '@/assets/components/layout';

export const Error = (props: { title: string }) => {
   return (
      <PageContentBlock title="Not Found" className="bg-stone-800 rounded-b-3xl">
         <section>
            <div className="mx-auto px-4 lg:px-8 max-w-screen 2xl:max-w-[120rem] 2xl:lg:px-24 pb-32 lg:pb-48 pt-12">
               <h1 className="text-5xl sm:text-4xl md:text-5xl lg:text-6xl text-stone-100 font-medium lg:w-1/3">Error 404</h1>
            </div>
         </section>
         <section>
            <div className="mx-auto px-4 lg:px-8 max-w-screen 2xl:max-w-[120rem] 2xl:lg:px-24 pb-8">
               <p className="text-base text-stone-300 text-balance max-w-xl">
                  Sorry, the page you are looking for does not exist. It might have been removed, renamed, or is temporarily unavailable.
               </p>
               <div className="mt-8 flex flex-wrap gap-4">
                  <Link
                     to="/"
                     className="shadow flex transition rounded-full items-center duration-300 focus:outline-2 focus:ring-none ring-offset-stone-950 focus:outline-offset-2 group relative text-black bg-white hover:text-accent-600 focus:outline-accent-600 h-10 px-6 py-3 text-base font-medium pr-10">
                     <span className="z-10 pr-2">Go back home</span>
                     <div className="absolute right-1 inline-flex h-8 w-8 items-center justify-end rounded-full bg-stone-300 transition-[width] group-hover:w-[calc(100%-8px)]">
                        <div className="relative h-8 w-8 flex items-center justify-center">
                           <svg className="size-4" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                              <path
                                 d="M8.14645 3.14645C8.34171 2.95118 8.65829 2.95118 8.85355 3.14645L12.8536 7.14645C13.0488 7.34171 13.0488 7.65829 12.8536 7.85355L8.85355 11.8536C8.65829 12.0488 8.34171 12.0488 8.14645 11.8536C7.95118 11.6583 7.95118 11.3417 8.14645 11.1464L11.2929 8H2.5C2.22386 8 2 7.77614 2 7.5C2 7.22386 2.22386 7 2.5 7H11.2929L8.14645 3.85355C7.95118 3.65829 7.95118 3.34171 8.14645 3.14645Z"
                                 fill="currentColor"
                                 fill-rule="evenodd"
                                 clip-rule="evenodd"></path>
                           </svg>
                        </div>
                     </div>
                  </Link>
               </div>
            </div>
         </section>
      </PageContentBlock>
   );
};
