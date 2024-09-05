import { Link, usePage } from "@inertiajs/react";

export default function Layout(props: { children: React.ReactNode }) {
	const { url } = usePage();

	return (
		<>
			<header className="border-b-4 border-b-purple-600">
				<div className="flex items-center py-4 container  max-w-[1280px] mx-auto space-x-8">
					<div className="font-medium">Go + Inertia</div>
					<nav className="flex space-x-4">
						<Link href="/" className={url == "/" ? "font-semibold" : ""}>
							Home
						</Link>
						<Link
							href="/countries"
							className={url == "/countries" ? "font-semibold" : ""}
						>
							Countries
						</Link>
					</nav>
				</div>
			</header>
			<main className="container max-w-[1280px] mx-auto">{props.children}</main>
		</>
	);
}
