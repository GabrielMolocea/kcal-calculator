var targetCalories = 0;
var totalKcal = 0;

// Function to fetch and display target calories
async function fetchTargetCalories() {
    try {
        const response = await fetch('http://localhost:8080/targetkcal');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const targetKcal = await response.json();
        const targetElement = document.getElementById('target');
        targetElement.textContent = `Target Calories: ${targetKcal.targetKcal}`;
    } catch (error) {
        console.error('Error fetching target calories:', error);
    }
}

// Function to fetch and display foods
async function fetchFoods() {
    try {
        const response = await fetch('http://localhost:8080/foods');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const foods = await response.json();
        const foodList = document.getElementById('food-list');
        const nameSelect = document.getElementById('name');
        foodList.innerHTML = '';
        nameSelect.innerHTML = '<option value="">Select a food</option>';
        foods.forEach(food => {
            const li = document.createElement('li');
            li.textContent = `${food.name} - ${food.calories} calories`;
            foodList.appendChild(li);

            const option = document.createElement('option');
            option.value = food.name;
            option.textContent = food.name;
            option.dataset.calories = food.calories;
            nameSelect.appendChild(option);
        });
    } catch (error) {
        console.error('Error fetching foods:', error);
    }
}
        // Function to add a new food
        async function addFood(event) {
            event.preventDefault();
            const name = document.getElementById('name').value;
            const calories = parseInt(document.getElementById('calories').value, 10);
        
            try {
                const response = await fetch('http://localhost:8080/newfood', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ name, calories })
                });
        
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
        
                fetchFoods();
                document.getElementById('food-form').reset();
                totalKcal = targetCalories - calories;
                const totalElement = document.getElementById('total');
                totalElement.textContent = `You have eat today this much calories: ${totalKcal}`;
            } catch (error) {
                console.error('Error adding food:', error);
                alert('Failed to add food');
            }
    } 

    async function resetCalories() {
        try {
            const response = await fetch('http://localhost:8080/resttargetkcal', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
        });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const targetKcal = await response.json();            
            const targetElement = document.getElementById('target');
            targetElement.textContent = `Target Calories: ${targetKcal.target}`;

       }
       catch (error) {
           console.error('Error resetting calories:', error);
       }
    }

    // Event listener to update calories input when a food is selected
    document.getElementById('name').addEventListener('change', function() {
    const selectedOption = this.options[this.selectedIndex];
    const calories = selectedOption.dataset.calories || '';
    document.getElementById('calories').value = calories;
    });
        
        // Fetch foods when the page loads
        document.addEventListener('DOMContentLoaded', () => {
            fetchTargetCalories();
            fetchFoods()});
        
        // Add event listener to the form
        document.getElementById('food-form').addEventListener('submit', addFood);
        document.getElementById('reset').addEventListener('click', resetCalories);