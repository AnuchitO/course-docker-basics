<script setup lang="ts">
withDefaults(defineProps<{ name?: string, open?: boolean, width?: number }>(), { open: true, width: 480 })
</script>

<template>
  <!-- Shipping container traced from a reference photo: perspective view, trapezoid door end,
       perspective-correct ribs that narrow toward the far end, ribbed roof, lock rods, corner castings.
       `open` slides the two doors apart along the door face; the frame clips them to reveal the app. Shapes only (no <text>): slide CSS can resize SVG text. -->
  <div class="cbig" :style="{ width: width + 'px' }">
    <svg viewBox="0 0 800 580" class="w-full" xmlns="http://www.w3.org/2000/svg">
      <ellipse cx="470" cy="548" rx="330" ry="22" fill="#000" opacity="0.14" />
      <polygon points="20.0,92.0 225.0,160.0 775.0,45.0 555.0,12.0" fill="#1b9fe3" />
      <line x1="42.8" y1="99.6" x2="579.4" y2="15.7" stroke="#4dbcf0" stroke-width="2.2" opacity="0.8" />
      <line x1="65.6" y1="107.1" x2="603.9" y2="19.3" stroke="#0f86c9" stroke-width="2.2" opacity="0.8" />
      <line x1="88.3" y1="114.7" x2="628.3" y2="23.0" stroke="#4dbcf0" stroke-width="2.2" opacity="0.8" />
      <line x1="111.1" y1="122.2" x2="652.8" y2="26.7" stroke="#0f86c9" stroke-width="2.2" opacity="0.8" />
      <line x1="133.9" y1="129.8" x2="677.2" y2="30.3" stroke="#4dbcf0" stroke-width="2.2" opacity="0.8" />
      <line x1="156.7" y1="137.3" x2="701.7" y2="34.0" stroke="#0f86c9" stroke-width="2.2" opacity="0.8" />
      <line x1="179.4" y1="144.9" x2="726.1" y2="37.7" stroke="#4dbcf0" stroke-width="2.2" opacity="0.8" />
      <line x1="202.2" y1="152.4" x2="750.6" y2="41.3" stroke="#0f86c9" stroke-width="2.2" opacity="0.8" />
      <polygon points="20.0,92.0 225.0,160.0 775.0,45.0 555.0,12.0" fill="none" stroke="#0b6aa8" stroke-width="3" stroke-linejoin="round" />
      <polygon points="225.0,160.0 775.0,45.0 775.0,308.0 228.0,540.0" fill="#0a7bc2" />
      <polygon points="225.2,188.5 238.7,185.5 241.2,509.8 227.8,515.3" fill="#1a9de0" />
      <polygon points="238.7,185.5 247.6,183.5 250.1,506.2 241.2,509.8" fill="#0b79bf" />
      <polygon points="247.6,183.5 255.0,181.8 257.4,503.2 250.1,506.2" fill="#075d98" />
      <polygon points="255.0,181.8 268.0,178.9 270.4,497.8 257.4,503.2" fill="#1a9de0" />
      <polygon points="268.0,178.9 276.7,176.9 279.0,494.3 270.4,497.8" fill="#0b79bf" />
      <polygon points="276.7,176.9 283.8,175.3 286.1,491.4 279.0,494.3" fill="#075d98" />
      <polygon points="283.8,175.3 296.4,172.5 298.7,486.2 286.1,491.4" fill="#1a9de0" />
      <polygon points="296.4,172.5 304.8,170.6 307.0,482.8 298.7,486.2" fill="#0b79bf" />
      <polygon points="304.8,170.6 311.6,169.0 313.8,480.0 307.0,482.8" fill="#075d98" />
      <polygon points="311.6,169.0 323.9,166.3 326.0,475.0 313.8,480.0" fill="#1a9de0" />
      <polygon points="323.9,166.3 332.0,164.5 334.1,471.7 326.0,475.0" fill="#0b79bf" />
      <polygon points="332.0,164.5 338.7,163.0 340.7,469.0 334.1,471.7" fill="#075d98" />
      <polygon points="338.7,163.0 350.6,160.3 352.6,464.1 340.7,469.0" fill="#1a9de0" />
      <polygon points="350.6,160.3 358.4,158.5 360.3,460.9 352.6,464.1" fill="#0b79bf" />
      <polygon points="358.4,158.5 364.9,157.1 366.8,458.3 360.3,460.9" fill="#075d98" />
      <polygon points="364.9,157.1 376.4,154.5 378.3,453.6 366.8,458.3" fill="#1a9de0" />
      <polygon points="376.4,154.5 384.0,152.8 385.8,450.5 378.3,453.6" fill="#0b79bf" />
      <polygon points="384.0,152.8 390.3,151.3 392.1,447.9 385.8,450.5" fill="#075d98" />
      <polygon points="390.3,151.3 401.5,148.8 403.2,443.4 392.1,447.9" fill="#1a9de0" />
      <polygon points="401.5,148.8 408.8,147.2 410.6,440.4 403.2,443.4" fill="#0b79bf" />
      <polygon points="408.8,147.2 414.9,145.8 416.6,437.9 410.6,440.4" fill="#075d98" />
      <polygon points="414.9,145.8 425.8,143.3 427.4,433.4 416.6,437.9" fill="#1a9de0" />
      <polygon points="425.8,143.3 432.9,141.7 434.5,430.5 427.4,433.4" fill="#0b79bf" />
      <polygon points="432.9,141.7 438.9,140.4 440.4,428.1 434.5,430.5" fill="#075d98" />
      <polygon points="438.9,140.4 449.4,138.0 450.9,423.8 440.4,428.1" fill="#1a9de0" />
      <polygon points="449.4,138.0 456.4,136.5 457.9,421.0 450.9,423.8" fill="#0b79bf" />
      <polygon points="456.4,136.5 462.1,135.2 463.6,418.6 457.9,421.0" fill="#075d98" />
      <polygon points="462.1,135.2 472.3,132.9 473.8,414.4 463.6,418.6" fill="#1a9de0" />
      <polygon points="472.3,132.9 479.1,131.3 480.5,411.7 473.8,414.4" fill="#0b79bf" />
      <polygon points="479.1,131.3 484.7,130.1 486.1,409.4 480.5,411.7" fill="#075d98" />
      <polygon points="484.7,130.1 494.6,127.8 496.0,405.3 486.1,409.4" fill="#1a9de0" />
      <polygon points="494.6,127.8 501.2,126.4 502.5,402.7 496.0,405.3" fill="#0b79bf" />
      <polygon points="501.2,126.4 506.6,125.1 507.9,400.4 502.5,402.7" fill="#075d98" />
      <polygon points="506.6,125.1 516.3,123.0 517.5,396.5 507.9,400.4" fill="#1a9de0" />
      <polygon points="516.3,123.0 522.7,121.5 523.9,393.9 517.5,396.5" fill="#0b79bf" />
      <polygon points="522.7,121.5 528.0,120.3 529.1,391.7 523.9,393.9" fill="#075d98" />
      <polygon points="528.0,120.3 537.4,118.2 538.5,387.9 529.1,391.7" fill="#1a9de0" />
      <polygon points="537.4,118.2 543.6,116.8 544.7,385.3 538.5,387.9" fill="#0b79bf" />
      <polygon points="543.6,116.8 548.8,115.7 549.8,383.2 544.7,385.3" fill="#075d98" />
      <polygon points="548.8,115.7 557.9,113.6 558.9,379.5 549.8,383.2" fill="#1a9de0" />
      <polygon points="557.9,113.6 564.0,112.2 565.0,377.0 558.9,379.5" fill="#0b79bf" />
      <polygon points="564.0,112.2 569.0,111.1 569.9,375.0 565.0,377.0" fill="#075d98" />
      <polygon points="569.0,111.1 577.9,109.1 578.8,371.4 569.9,375.0" fill="#1a9de0" />
      <polygon points="577.9,109.1 583.8,107.8 584.7,369.0 578.8,371.4" fill="#0b79bf" />
      <polygon points="583.8,107.8 588.6,106.7 589.5,367.0 584.7,369.0" fill="#075d98" />
      <polygon points="588.6,106.7 597.3,104.7 598.2,363.4 589.5,367.0" fill="#1a9de0" />
      <polygon points="597.3,104.7 603.1,103.4 603.9,361.1 598.2,363.4" fill="#0b79bf" />
      <polygon points="603.1,103.4 607.8,102.4 608.6,359.2 603.9,361.1" fill="#075d98" />
      <polygon points="607.8,102.4 616.3,100.5 617.0,355.7 608.6,359.2" fill="#1a9de0" />
      <polygon points="616.3,100.5 621.8,99.2 622.6,353.4 617.0,355.7" fill="#0b79bf" />
      <polygon points="621.8,99.2 626.5,98.2 627.2,351.5 622.6,353.4" fill="#075d98" />
      <polygon points="626.5,98.2 634.7,96.3 635.4,348.2 627.2,351.5" fill="#1a9de0" />
      <polygon points="634.7,96.3 640.1,95.1 640.8,345.9 635.4,348.2" fill="#0b79bf" />
      <polygon points="640.1,95.1 644.6,94.1 645.3,344.1 640.8,345.9" fill="#075d98" />
      <polygon points="644.6,94.1 652.7,92.3 653.3,340.8 645.3,344.1" fill="#1a9de0" />
      <polygon points="652.7,92.3 658.0,91.1 658.5,338.7 653.3,340.8" fill="#0b79bf" />
      <polygon points="658.0,91.1 662.4,90.1 662.9,336.9 658.5,338.7" fill="#075d98" />
      <polygon points="662.4,90.1 670.2,88.3 670.7,333.7 662.9,336.9" fill="#1a9de0" />
      <polygon points="670.2,88.3 675.4,87.2 675.8,331.6 670.7,333.7" fill="#0b79bf" />
      <polygon points="675.4,87.2 679.7,86.2 680.1,329.8 675.8,331.6" fill="#075d98" />
      <polygon points="679.7,86.2 687.3,84.5 687.7,326.7 680.1,329.8" fill="#1a9de0" />
      <polygon points="687.3,84.5 692.4,83.3 692.7,324.6 687.7,326.7" fill="#0b79bf" />
      <polygon points="692.4,83.3 696.5,82.4 696.9,322.9 692.7,324.6" fill="#075d98" />
      <polygon points="696.5,82.4 704.0,80.7 704.3,319.9 696.9,322.9" fill="#1a9de0" />
      <polygon points="704.0,80.7 708.9,79.6 709.2,317.9 704.3,319.9" fill="#0b79bf" />
      <polygon points="708.9,79.6 713.0,78.7 713.3,316.2 709.2,317.9" fill="#075d98" />
      <polygon points="713.0,78.7 720.3,77.0 720.5,313.2 713.3,316.2" fill="#1a9de0" />
      <polygon points="720.3,77.0 725.1,76.0 725.3,311.3 720.5,313.2" fill="#0b79bf" />
      <polygon points="725.1,76.0 729.0,75.1 729.3,309.7 725.3,311.3" fill="#075d98" />
      <polygon points="729.0,75.1 736.1,73.5 736.3,306.8 729.3,309.7" fill="#1a9de0" />
      <polygon points="736.1,73.5 740.8,72.4 741.0,304.8 736.3,306.8" fill="#0b79bf" />
      <polygon points="740.8,72.4 744.7,71.5 744.9,303.3 741.0,304.8" fill="#075d98" />
      <polygon points="744.7,71.5 751.7,70.0 751.8,300.4 744.9,303.3" fill="#1a9de0" />
      <polygon points="751.7,70.0 756.2,68.9 756.3,298.6 751.8,300.4" fill="#0b79bf" />
      <polygon points="756.2,68.9 760.0,68.1 760.1,297.0 756.3,298.6" fill="#075d98" />
      <polygon points="760.0,68.1 766.8,66.6 766.9,294.2 760.1,297.0" fill="#1a9de0" />
      <polygon points="766.8,66.6 771.3,65.6 771.3,292.4 766.9,294.2" fill="#0b79bf" />
      <polygon points="771.3,65.6 775.0,64.7 775.0,290.9 771.3,292.4" fill="#075d98" />
      <polygon points="225.0,160.0 775.0,45.0 775.0,64.7 225.2,188.5" fill="#2aa9ea" />
      <polygon points="227.8,515.3 775.0,290.9 775.0,308.0 228.0,540.0" fill="#0a68a8" />
      <polygon points="225.0,160.0 775.0,45.0 775.0,308.0 228.0,540.0" fill="none" stroke="#0a5d96" stroke-width="3" stroke-linejoin="round" />
      <polygon points="435.9,445.8 477.7,428.1 477.7,444.1 435.9,461.8" fill="#0a3558" opacity="0.85" />
      <polygon points="528.4,406.6 565.6,390.8 565.6,406.8 528.4,422.6" fill="#0a3558" opacity="0.85" />
      <polygon points="20.0,92.0 225.0,160.0 228.0,540.0 22.0,405.0" fill="#1590d6" />
      <polygon points="28.3,108.9 216.9,174.3 219.6,517.6 30.1,396.2" fill="#0a2d4d" />
      <g class="app" :class="{ on: open }"><polygon points="40.8,150.0 205.0,212.9 206.9,466.8 42.3,367.3" fill="#ffffff" /><polygon points="40.8,150.0 205.0,212.9 205.4,265.2 41.1,194.7" fill="#dbe4ee" /><circle cx="55.3" cy="178.2" r="5.5" fill="#ef4444" /><circle cx="69.7" cy="184.0" r="5.5" fill="#f59e0b" /><circle cx="84.1" cy="189.8" r="5.5" fill="#22c55e" /><circle cx="78.7" cy="296.8" r="21" fill="#2496ed" /><polygon points="111.3,276.0 189.4,313.3 189.6,339.1 111.5,300.0" fill="#94a3b8" /><polygon points="111.7,327.4 161.0,353.4 161.2,378.5 111.8,351.4" fill="#cbd5e1" /></g>
      <defs><clipPath id="cbig-face"><polygon points="28.3,108.9 216.9,174.3 219.6,517.6 30.1,396.2" /></clipPath></defs><g class="doors" clip-path="url(#cbig-face)" :class="{ open }"><g class="dl"><polygon points="43.3,208.4 55.6,213.9 55.8,252.8 43.5,246.9" fill="#e2e8ee" /><polygon points="30.4,119.1 121.7,151.6 123.8,445.9 32.1,388.0" fill="#1a9be0" stroke="#0a3d66" stroke-width="2.5" /><line x1="65.2" y1="126.6" x2="67.2" y2="415.0" stroke="#cfd8e1" stroke-width="5" stroke-linecap="round" /><rect x="59.3" y="136.7" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="60.5" y="307.2" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="60.8" y="359.6" width="16" height="6" rx="2" fill="#aeb9c4" /><line x1="102.1" y1="139.6" x2="104.3" y2="438.6" stroke="#cfd8e1" stroke-width="5" stroke-linecap="round" /><rect x="96.2" y="150.2" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="97.5" y="326.9" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="97.9" y="381.2" width="16" height="6" rx="2" fill="#aeb9c4" /></g><g class="dr"><polygon points="123.7,152.4 215.0,184.8 217.5,505.0 125.8,447.2" fill="#1a9be0" stroke="#0a3d66" stroke-width="2.5" /><line x1="137.0" y1="151.8" x2="139.3" y2="460.9" stroke="#cfd8e1" stroke-width="5" stroke-linecap="round" /><rect x="131.1" y="162.9" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="132.4" y="345.5" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="132.9" y="401.7" width="16" height="6" rx="2" fill="#aeb9c4" /><line x1="180.1" y1="167.0" x2="182.5" y2="488.4" stroke="#cfd8e1" stroke-width="5" stroke-linecap="round" /><rect x="174.2" y="178.6" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="175.6" y="368.5" width="16" height="6" rx="2" fill="#aeb9c4" /><rect x="176.1" y="426.9" width="16" height="6" rx="2" fill="#aeb9c4" /></g></g>
      <polygon points="8.9,75.8 31.0,82.8 31.2,108.5 9.0,100.9" fill="#0a3a66" />
      <polygon points="213.8,141.1 235.9,148.1 236.2,179.2 214.0,171.6" fill="#0a3a66" />
      <polygon points="10.8,385.2 33.0,399.5 33.2,425.1 11.0,410.2" fill="#0a3a66" />
      <polygon points="216.8,517.5 239.0,531.8 239.2,562.8 217.0,548.0" fill="#0a3a66" />
      <rect x="761.0" y="35.0" width="22" height="20" rx="3" fill="#0a3a66" />
      <rect x="761.0" y="298.0" width="22" height="20" rx="3" fill="#0a3a66" />
    </svg>
    <div v-if="name" class="label"><span class="dot" :class="{ on: open }" />{{ name }} <span class="state">{{ open ? 'running' : 'created' }}</span></div>
  </div>
</template>

<style scoped>
.cbig { display: flex; flex-direction: column; align-items: center; }
.doors .dl, .doors .dr { transition: transform 0.9s cubic-bezier(.65,0,.25,1); transform-box: view-box; transform-origin: 0 0; }
.doors.open .dl { transform: matrix(1.00052,0.03477,-0.00146,0.90235,-94.18,-23.02); }
.doors.open .dr { transform: matrix(0.99953,-0.03172,0.00133,1.08909,94.19,23.94); }
.app { opacity: 0; transition: opacity 0.5s 0.25s; }
.app.on { opacity: 1; }
.label { margin-top: 0.3rem; font-weight: 700; display: flex; align-items: center; gap: 0.5rem; }
.state { font-weight: 600; opacity: 0.6; font-size: 0.85em; }
.dot { width: 0.7rem; height: 0.7rem; border-radius: 50%; background: #94a3b8; transition: background 0.4s; }
.dot.on { background: #22c55e; box-shadow: 0 0 0 4px rgba(34,197,94,.25); }
</style>
