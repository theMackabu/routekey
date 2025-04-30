import { useEffect } from 'react';

export const PageContentBlock = ({ children, title, ...props }) => {
   useEffect(() => {
      document.title = title ? title + ' | RouteKey' : 'RouteKey';
   }, [title]);

   return <div {...props}>{children}</div>;
};

export const Layout = ({ header, footer, title, children }) => (
   <div className="flex flex-col min-h-screen">
      {header ? <header className="w-full mt-3">{header}</header> : null}
      <main className="flex-grow flex items-center justify-center -mt-16">
         <PageContentBlock title={title}>
            <section className="max-w-7xl mx-auto">{children}</section>
         </PageContentBlock>
      </main>
      {footer ? <footer className="w-full mb-3">{footer}</footer> : null}
   </div>
);
