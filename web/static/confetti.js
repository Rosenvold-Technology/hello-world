
document.addEventListener("DOMContentLoaded", () => {
  const partyBtn = document.getElementById("party");
  partyBtn.addEventListener("click", () => {
    import("https://cdn.skypack.dev/canvas-confetti").then((module) => {
      const confetti = module.default;
      confetti({
        particleCount: 200,
        spread: 120,
        origin: { y: 0.6 }
      });
    });
  });
});
