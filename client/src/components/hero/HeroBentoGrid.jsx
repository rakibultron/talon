import React from "react";

const HeroBentoGrid = () => {
  const gridItems = [
    // Full width on mobile, 6 cols and 2 rows on desktop
    {
      id: 1,
      text: "Item 1",
      className: "col-span-12 md:col-span-6 md:row-span-2",
    },
    // Full width on mobile, 6 cols on desktop
    { id: 2, text: "Item 2", className: "col-span-12 md:col-span-6" },
    // Half width on mobile, 3 cols on desktop
    { id: 3, text: "Item 3", className: "col-span-6 md:col-span-3" },
    // Half width on mobile, 3 cols on desktop
    { id: 4, text: "Item 4", className: "col-span-6 md:col-span-3" },
  ];

  return (
    <div className="p-4">
      <div className="grid grid-cols-12 gap-4 auto-rows-[150px]">
        {gridItems.map((item) => (
          <div
            key={item.id}
            className={`bg-gray-200 rounded-md shadow-md flex items-center justify-center text-center p-4 ${item.className}`}
          >
            {item.text}
          </div>
        ))}
      </div>
    </div>
  );
};

export default HeroBentoGrid;
