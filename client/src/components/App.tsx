import React, { useEffect } from 'react';

import { GlobalStyles } from '@/assets/styles';
import { ScrollToTop } from '@/components/elements/generic';
import { Home, NotFound } from '@/components/pages/generic';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

const Page = (props: { component: any; title?: any }) => {
	return <props.component title={props.title} />;
};

const App = () => {
	return (
		<BrowserRouter>
			<GlobalStyles />
			<ScrollToTop>
				<Routes>
					<Route path='*' element={<NotFound title='Whoops, cant find' />} />
					<Route path='/' element={<Page component={Home} />} />
					<Route path='/app/admin' element={<div>test</div>} />
				</Routes>
			</ScrollToTop>
		</BrowserRouter>
	);
};

export default App;
