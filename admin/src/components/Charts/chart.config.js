export const chartColors = {
  default: {
    primary: "#00D1B2",
    info: "#209CEE",
    danger: "#FF3860",
  },
};

const randomChartData = (n) => {
  const data = [];

  for (let i = 0; i <= n; i++) {
    data.push(Math.round(Math.random() * 3));
  }

  return data;
};

const datasetObject = (color, points) => {
  return {
    fill: false,
    borderColor: chartColors.default[color],
    borderWidth: 2,
    borderDash: [],
    borderDashOffset: 0.0,
    pointBackgroundColor: chartColors.default[color],
    pointBorderColor: "rgba(255,255,255,0)",
    pointHoverBackgroundColor: chartColors.default[color],
    pointBorderWidth: 20,
    pointHoverRadius: 4,
    pointHoverBorderWidth: 15,

    pointRadius: 4,
    data: randomChartData(points),
    tension: 0.5,
    cubicInterpolationMode: "default",
  };
};

function formatDate(date) {
  const day = date.getDate().toString().padStart(2, '0');
  const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  const month = monthNames[date.getMonth()];
  return `${day}-${month}`;
}


export const sampleChartData = (points = 6) => {
  const labels = [];


  for (let i = points; i >= 0; i--) {
    const currentDate = new Date();
    currentDate.setDate(currentDate.getDate() - i);
    const formattedDate = formatDate(currentDate);
//    console.log(formattedDate);
    labels.push(formattedDate);
  }

  
/*   for (let i = 1; i <= points; i++) {
    labels.push(`0${i}`);
  }
 */
  return {
    labels,
    datasets: [
      datasetObject("primary", points),
      datasetObject("info", points),
    ],
  };
};
