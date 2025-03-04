import { useState, useCallback } from "react";
import { Search, Plus } from "lucide-react";

interface CaloriesFormProps {
	setTotalKcal: (kcal: number) => void;
	getTargetCalories: () => void;
	foodList: { name: string; calories: number }[];
}

const CaloriesForm: React.FC<CaloriesFormProps> = ({
	setTotalKcal,
	foodList,
	getTargetCalories,
}) => {
	const [searchTerm, setSearchTerm] = useState("");
	const [selectedFood, setSelectedFood] = useState<{
		name: string;
		calories: number;
	} | null>(null);
	const [isListVisible, setIsListVisible] = useState(false);
	const [showNewFoodForm, setShowNewFoodForm] = useState(false);
	const [newFood, setNewFood] = useState({ name: "", calories: "" });

	const handleSearchChange = useCallback(
		(event: React.ChangeEvent<HTMLInputElement>) => {
			setSearchTerm(event.target.value);
			setIsListVisible(true);
		},
		[]
	);

	const handleFoodSelect = useCallback(
		(food: { name: string; calories: number }) => {
			setSelectedFood(food);
			setSearchTerm(food.name);
			setIsListVisible(false);
			setTotalKcal(food.calories);
		},
		[]
	);

	const filteredFoods = foodList
		.filter((food: { name: string }) =>
			food.name.toLowerCase().includes(searchTerm.toLowerCase())
		)
		.slice(0, 10);

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		if (!selectedFood) return;

		try {
			await fetch("http://localhost:8080/deductkcal", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(selectedFood),
			});
			setSelectedFood(null);
			setSearchTerm("");
		} catch (error) {
			console.error("Failed to add food:", error);
		} finally {
			getTargetCalories();
		}
	};

	const handleNewFoodSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		if (!newFood.name || !newFood.calories) return;

		try {
			await fetch("http://localhost:8080/newfood", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					name: newFood.name,
					calories: parseInt(newFood.calories),
				}),
			});
			setNewFood({ name: "", calories: "" });
			setShowNewFoodForm(false);
		} catch (error) {
			console.error("Failed to add new food:", error);
		}
	};

	return (
		<div className="w-full max-w-md mx-auto p-4">
			<div className="space-y-4">
				<div className="flex justify-between items-center">
					<div className="text-xl font-semibold">Add Food</div>
					<button
						onClick={() => setShowNewFoodForm(!showNewFoodForm)}
						className="text-blue-500 hover:text-blue-600 flex items-center gap-1"
					>
						<Plus size={20} />
						{showNewFoodForm ? "Cancel" : "Add New Food"}
					</button>
				</div>

				{showNewFoodForm ? (
					<form
						onSubmit={handleNewFoodSubmit}
						className="space-y-4 p-4 bg-gray-50 rounded-lg"
					>
						<div className="space-y-2">
							<input
								type="text"
								value={newFood.name}
								onChange={(e) =>
									setNewFood((prev) => ({ ...prev, name: e.target.value }))
								}
								placeholder="Food name"
								className="w-full p-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
							/>
							<input
								type="number"
								value={newFood.calories}
								onChange={(e) =>
									setNewFood((prev) => ({ ...prev, calories: e.target.value }))
								}
								placeholder="Calories"
								className="w-full p-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
							/>
						</div>
						<button
							type="submit"
							disabled={!newFood.name || !newFood.calories}
							className="w-full bg-blue-500 text-white p-2 rounded-lg hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed"
						>
							Add New Food
						</button>
					</form>
				) : (
					<form onSubmit={handleSubmit} className="space-y-4">
						<div className="relative">
							<div className="relative">
								<input
									type="text"
									value={searchTerm}
									onChange={handleSearchChange}
									placeholder="Search food..."
									className="w-full p-2 pr-10 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
								/>
								<Search className="absolute right-3 top-2.5 text-gray-400" size={20} />
							</div>

							{isListVisible && searchTerm && (
								<ul className="absolute z-10 w-full mt-1 bg-white border rounded-lg shadow-lg max-h-60 overflow-y-auto">
									{filteredFoods.length > 0 ? (
										filteredFoods.map((food: { id?: any; name: any; calories: any }) => (
											<ul
												key={food.id}
												onClick={() => handleFoodSelect(food)}
												className="px-4 py-2 hover:bg-gray-100 cursor-pointer flex justify-between items-center"
											>
												<span>{food.name}</span>
												<span className="text-gray-500">{food.calories} cal</span>
											</ul>
										))
									) : (
										<li className="px-4 py-2 text-gray-500">No foods found</li>
									)}
								</ul>
							)}
						</div>

						<div className="flex gap-4">
							<input
								type="number"
								value={selectedFood?.calories || ""}
								readOnly
								placeholder="Calories"
								className="w-1/2 p-2 border rounded-lg bg-gray-50"
							/>

							<button
								type="submit"
								disabled={!selectedFood}
								className="w-1/2 bg-blue-500 text-white p-2 rounded-lg hover:bg-blue-600 disabled:bg-gray-300 disabled:cursor-not-allowed"
							>
								Add Food
							</button>
						</div>
					</form>
				)}
			</div>
		</div>
	);
};

export default CaloriesForm;
