<!DOCTYPE html />
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Tailwind CSS Responsive Sidebar UI </title>
		<script src="https://cdn.tailwindcss.com/"></script>
	</head>

	<body>
		<nav class="fixed z-30 w-full bg-white border-b-2 border-indigo-600">
			<div class="px-6 py-3">
				<div class="flex items-center justify-between">
					<div class="flex items-center justify-start">
						<button class="p-2 text-gray-600 rounded cursor-pointer lg:hidden ">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="w-6 h-6"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M4 6h16M4 12h16M4 18h16"
								/>
							</svg>
						</button>
						<a href="#" class="flex items-center text-xl font-bold">
							<span class="text-blue-800">Logo</span>
						</a>
					</div>
					<div class="flex items-center">
						<div class="hidden mr-6 lg:block">
							<form action="#">
								<label class="sr-only">Search</label>
								<div class="relative mt-1 lg:w-64">
									<div class="absolute inset-y-0 left-0 flex items-center pl-3">
										<svg
											class="w-5 h-5 text-gray-500"
											fill="currentColor"
											viewBox="0 0 20 20"
											xmlns="http://www.w3.org/2000/svg"
										>
											<path
												fill-rule="evenodd"
												d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
												clip-rule="evenodd"
											/>
										</svg>
									</div>
									<input
										type="text"
										name="name"
										class=" border  text-gray-900 sm:text-sm rounded-lg focus:outline-none focus:ring-1 block w-full pl-10 p-2.5"
										placeholder="Search"
									/>
								</div>
							</form>
						</div>
						<div>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="w-6 h-6"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
								/>
							</svg>
						</div>

						<div class="relative inline-block ">
							<!-- Dropdown toggle button -->
							<button
								class="relative flex items-center p-2 text-sm text-gray-600 bg-white border border-transparent rounded-md focus:border-blue-500 focus:ring-opacity-40 dark:focus:ring-opacity-40 focus:ring-blue-300 dark:focus:ring-blue-400 focus:ring dark:text-white dark:bg-gray-800 focus:outline-none"
							>
								<span class="mx-1">Jane Doe</span>
								<svg
									class="w-5 h-5 mx-1"
									viewBox="0 0 24 24"
									fill="none"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										d="M12 15.713L18.01 9.70299L16.597 8.28799L12 12.888L7.40399 8.28799L5.98999 9.70199L12 15.713Z"
										fill="currentColor"
									/>
								</svg>
							</button>

							<!-- Dropdown menu -->
							<div class="absolute right-0 z-20 w-56 mt-2 overflow-hidden bg-white rounded-md" />
						</div>
					</div>
				</div>
			</div>
		</nav>
		<div class="pt-12 lg:flex">
			<div
				class="flex flex-col w-full px-4 py-8 overflow-y-auto border-b lg:border-r lg:h-screen lg:w-64"
			>
				<div class="flex flex-col justify-between mt-6">
					<aside>
						<ul>
							<li>
								<a
									class="flex items-center px-4 py-2 text-gray-700 bg-gray-100 rounded-md "
									href="#"
								>
									<svg
										xmlns="http://www.w3.org/2000/svg"
										class="w-6 h-6"
										fill="none"
										viewBox="0 0 24 24"
										stroke="currentColor"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											stroke-width="2"
											d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
										/>
									</svg>

									<span class="mx-4 font-medium">Dashboard</span>
								</a>
							</li>

							<li>
								<a
									class="flex items-center px-4 py-2 mt-5 text-gray-600 rounded-md hover:bg-gray-200"
									href="#"
								>
									<svg
										xmlns="http://www.w3.org/2000/svg"
										class="w-6 h-6"
										fill="none"
										viewBox="0 0 24 24"
										stroke="currentColor"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											stroke-width="2"
											d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
										/>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											stroke-width="2"
											d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
										/>
									</svg>

									<span class="mx-4 font-medium">Settings</span>
								</a>
							</li>
						</ul>
					</aside>
				</div>
			</div>
			<div class="w-full h-full p-4 m-8 overflow-y-auto">
				<div class="flex items-center justify-center p-16 mr-8 border-4 border-dotted lg:p-40">
					Content...
				</div>
			</div>
		</div></body
	>
</html>
