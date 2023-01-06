import React, { useState, useEffect, Fragment } from 'react';

import tw from 'twin.macro';
import { Link } from 'react-router-dom';
import { compactLink } from '@/api/link';
import { Dialog } from '@headlessui/react';
import { PageContentBlock, Transition } from '@/components/elements/generic';
import { MailIcon, HeartIcon, KeyIcon, ArrowsExpandIcon, QrcodeIcon } from '@heroicons/react/solid';

const Home = (props: { title: any }) => {
	const [show, setShow] = useState(false);
	const [view, setView] = useState(false);
	const [prompt, setPrompt] = useState('');
	const [linkName, setLinkName] = useState('');
	const [linkUUID, setLinkUUID] = useState('');
	const [formData, setFormData] = useState({ target: '', expire_in: '300' });

	const handleChange = (event: any) => {
		setFormData({
			...formData,
			[event.target.name]: event.target.value,
		});
	};

	const handleSubmit = (event: any) => {
		event.preventDefault();
		if (/^(ftp|http|https):\/\/[^ "]+$/.test(formData.target)) {
			compactLink(formData.target, formData.expire_in)
				.then((data: any) => {
					setLinkName(data.link);
					setLinkUUID(data.id);
					setPrompt('');
				})
				.catch((err) => {
					compactLink(formData.target, formData.expire_in)
						.then((data: any) => {
							setLinkName(data.link);
							setLinkUUID(data.id);
							setPrompt('');
						})
						.catch((err) => {
							console.log(err);
							setPrompt('');
						});
					setPrompt('');
				});
		} else {
			setPrompt('Please enter a valid url.');
		}
	};

	const transitionPropsOverlay = {
		enter: tw`ease-out duration-300`,
		enterFrom: tw`opacity-0`,
		enterTo: tw`opacity-100`,
		leave: tw`ease-in duration-200`,
		leaveFrom: tw`opacity-100`,
		leaveTo: tw`opacity-0`,
	};

	const transitionPropsModal = {
		enter: tw`ease-out duration-300`,
		enterFrom: tw`opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95`,
		enterTo: tw`opacity-100 translate-y-0 sm:scale-100`,
		leave: tw`ease-in duration-200`,
		leaveFrom: tw`opacity-100 translate-y-0 sm:scale-100`,
		leaveTo: tw`opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95`,
	};

	return (
		<PageContentBlock title={props.title}>
			<Transition show={show} as={Fragment}>
				<Dialog as='div' tw='fixed z-10 inset-0 overflow-y-auto' onClose={setShow}>
					<div tw='flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0'>
						<Transition.Child as={Fragment} {...transitionPropsOverlay}>
							<Dialog.Overlay tw='fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity' />
						</Transition.Child>
						<span tw='hidden sm:inline-block sm:align-middle sm:h-screen' aria-hidden='true'>
							&#8203;
						</span>
						<Transition.Child as={Fragment} {...transitionPropsModal}>
							<div tw='inline-block align-bottom bg-white rounded-lg px-4 p-5 text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-7xl'>
								<div tw='mt-3 text-center sm:mt-5 p-6'>
									<Dialog.Title as='h3' tw='text-[40px] leading-6 font-normal text-black'>
										rk.ausdk12.org/
									</Dialog.Title>
									<div tw='mt-2'>
										<p tw='p-4 text-[130px] text-black font-bold'>{linkName}</p>
									</div>
								</div>
							</div>
						</Transition.Child>
					</div>
				</Dialog>
			</Transition>
			<Transition show={view} as={Fragment}>
				<Dialog as='div' tw='fixed z-10 inset-0 overflow-y-auto' onClose={setView}>
					<div tw='flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0'>
						<Transition.Child as={Fragment} {...transitionPropsOverlay}>
							<Dialog.Overlay tw='fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity' />
						</Transition.Child>
						<span tw='hidden sm:inline-block sm:align-middle sm:h-screen' aria-hidden='true'>
							&#8203;
						</span>
						<Transition.Child as={Fragment} {...transitionPropsModal}>
							<div tw='inline-block align-bottom bg-white rounded-lg p-1 text-left overflow-hidden shadow-xl transform transition-all sm:align-middle sm:max-w-7xl'>
								<img tw='rounded-lg' src={`/api/v2/links/${linkUUID}/qrcode`} />
							</div>
						</Transition.Child>
					</div>
				</Dialog>
			</Transition>
			<div tw='mt-4 flex items-center justify-center text-gray-500 text-lg font-medium'>
				RouteKey: A url shortener built for the school on Key Route
			</div>
			<div tw='flex items-center justify-center'>
				<span tw='mt-1 border-b border-gray-300 w-[30rem]' />
			</div>
			<div tw='mt-16 grid place-items-center'>
				<form onSubmit={handleSubmit}>
					<div tw='bg-white px-14 py-3 rounded-md border border-gray-300 shadow-sm'>
						<span tw='text-[70px] font-medium text-gray-700 flex items-center justify-center'>
							RouteKey <span tw='mt-8 ml-1.5 text-[40px] text-rose-900'>v2</span>
						</span>
						<p tw='text-sm -mt-2 mb-5 text-gray-600'>Enter url to generate a routekey (automatically expires in 5 minutes by default)</p>
						<div tw='-mx-6 relative border border-gray-300 rounded-md px-3 py-2 hover:shadow-sm focus-within:ring-1 focus-within:ring-rose-800 focus-within:border-rose-800 transition'>
							<label htmlFor='target' tw='absolute -top-2 left-2 -mt-px inline-block px-1 bg-white text-xs font-medium text-gray-900'>
								Full URL
							</label>
							<input
								type='text'
								name='target'
								id='target'
								required
								value={formData.target}
								onChange={handleChange}
								tw='block w-full border-0 p-0 text-gray-900 placeholder-gray-500 focus:ring-0 sm:text-sm transition'
								placeholder='http://blog.interviewing.io/can-fake-names-create-bias/'
							/>
						</div>
						{prompt !== '' && <span tw='-ml-4 text-[13px] text-red-500 mt-0.5'>{prompt}</span>}
						<div tw='mt-5 -mx-6 relative border border-gray-300 rounded-md px-3 py-2 hover:shadow-sm focus-within:ring-1 focus-within:ring-rose-800 focus-within:border-rose-800 transition'>
							<label htmlFor='expire_in' tw='absolute -top-2 left-2 -mt-px inline-block px-1 bg-white text-xs font-medium text-gray-900'>
								Expires in
							</label>
							<select
								id='expire_in'
								name='expire_in'
								value={formData.expire_in}
								onChange={handleChange}
								tw='block w-full border-0 p-0 text-gray-700 focus:ring-0 sm:text-sm transition cursor-pointer'>
								<option value='300'>5 Minutes</option>
								<option value='900'>15 Minutes</option>
								<option value='1800'>30 Minutes</option>
								<option value='3600'>1 Hour</option>
							</select>
						</div>
						<div tw='-mx-6 mt-4 flex flex-col justify-center mb-4'>
							<button
								type='submit'
								tw='items-center py-2 border border-rose-900 hover:border-rose-800 hover:shadow text-lg font-semibold rounded-md text-white bg-rose-800 hover:bg-rose-700 transition'>
								Generate!
							</button>
						</div>
						<div tw='-mx-10 mt-0.5'>
							<div tw='grid place-items-center'>
								{linkName !== '' && (
									<Fragment>
										<span tw='text-[1.5rem] text-gray-800'>
											Your generated RouteKey is:
											<span tw='inline ml-1 text-[2rem] text-black'>
												{linkName}
												<button
													onClick={() => setShow(true)}
													type='button'
													tw='ml-2.5 inline p-1 rounded-md bg-rose-200 text-rose-800 hover:bg-rose-300 transition hover:shadow-sm hover:ring-1 hover:ring-rose-500'>
													<ArrowsExpandIcon tw='h-5 w-5' />
												</button>
												<button
													onClick={() => setView(true)}
													type='button'
													tw='ml-1.5 inline p-1 rounded-md bg-green-200 text-green-800 hover:bg-green-300 transition hover:shadow-sm hover:ring-1 hover:ring-green-500'>
													<QrcodeIcon tw='h-5 w-5' />
												</button>
											</span>
										</span>
										<p tw='text-lg block text-gray-700'>
											Go to:{' '}
											<a tw='text-blue-600' href={`https://rk.ausdk12.org/${linkName}`}>
												https://rk.ausdk12.org/{linkName}
											</a>{' '}
											to use
										</p>
									</Fragment>
								)}
							</div>
						</div>
					</div>
				</form>
				<div tw='w-[631.67px] mt-4 bg-white px-14 py-3 rounded-md border border-gray-300 shadow-sm'>
					<div tw='flow-root'>
						<div tw='-mx-[2.3rem] float-left mt-0.5'>
							<p tw='text-xl'>Have an ausdk12 staff account?</p>
							<p tw='text-gray-600'>Access the admin page and its associated features</p>
						</div>
						<button tw='mt-0.5 -mr-[2.4rem] float-right p-4 rounded-md bg-blue-200 text-blue-800 hover:bg-blue-300 transition hover:shadow-sm hover:ring-1 hover:ring-blue-500'>
							<KeyIcon tw='h-5 w-5' />
						</button>
					</div>
				</div>
			</div>
			<div tw='mt-8 mb-2 flex items-center justify-center'>
				<a
					href='https://github.com/theMackabu'
					target='_blank'
					rel='noopener noreferrer'
					className='group'
					tw='inline-flex text-sm font-medium text-gray-600 hover:text-gray-700 transition'>
					Made with <HeartIcon tw='mt-0.5 mx-1 text-red-400 h-4 w-4 group-hover:text-red-500 transition' /> in Albany, California
				</a>
			</div>
		</PageContentBlock>
	);
};

export default Home;
