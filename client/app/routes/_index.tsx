import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "Chess" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  return (
    <div className="flex justify-center items-center min-h-screen">
			<div className="w-[30vh] top-[10vh] h-[10vh] absolute flex justify-center items-center bg-black border-green-500 rounded-xl border-4 z-10">
				<p className="text-[8vh] font-bold z-10 text-green-500 font-mono">
					CHESS
				</p>
			</div>
			<img 
				className="w-full lg:w-1/3 filter blur-sm" 
				src="/board.png"
			/>
			<div className="top-1/3 h-[30vh] absolute flex-col justify-center items-center">
				<button className="w-[40vh] text-[3vh] flex border-4 rounded-xl bg-black border-green-500 font-mono justify-center items-center">
					<p>
						Start Game
					</p>
				</button>
				<div className="h-1/6"></div>
				<button className="w-[40vh] text-[3vh] block border-4 rounded-xl bg-black border-green-500 font-mono justify-center items-center">
					<textarea className="w-full text-[2vh] h-[4vh] border-2 rounded-xl bg-grey-500 border-green-500 justify-center">
						Enter a game code to join
					</textarea>
					<p>
						Join Game
					</p>
				</button>
			</div>
    </div>
  )
}