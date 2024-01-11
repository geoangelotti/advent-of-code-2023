(ns day-04.core
  (:require
   [clojure.set :as set]))

(defn get-sets [atoms]
  (->> atoms
       (map #(clojure.string/trim %))
       (remove empty?)
       (map #(Integer. %))
       (set)))

(defn extract-numbers [line]
  (->> line
       (#(clojure.string/split % #"\:"))
       (last)
       (#(clojure.string/split % #"\|"))
       (map #(clojure.string/split % #" "))
       (map get-sets)))

(defn evaluate-score [winners]
  (->> winners
       (reduce clojure.set/intersection)
       (count)
       (#(if (= % 0) 0 (Math/pow 2 (dec %))))
       (int)))

(defn lines [input] (clojure.string/split-lines input))

(defn process-part-1 [input]
  (->> input
       (lines)
       (map extract-numbers)
       (map evaluate-score)
       (reduce +)))

(defn process-part-2 [input] 0)