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

(defn process-part-1 [input]
  (->> input
       (clojure.string/split-lines)
       (map extract-numbers)
       (map evaluate-score)
       (reduce +)))


(defn process-part-2 [input]
  (let [lines (clojure.string/split-lines input)
        init-cards (reduce
                    (fn [cards n]
                      (assoc cards (inc n) 1))
                    {}
                    (range (count (clojure.string/split-lines input))))]
    (->> lines
         (reduce (fn [cards line]
                   (let [[winning-line my-number] (clojure.string/split line #"\|")
                         [card-line winning-numbers] (clojure.string/split winning-line #":")
                         card-number (parse-long (re-find #"\d+" card-line))
                         matches (count
                                  (set/intersection
                                   (set (re-seq #"\d+" winning-numbers))
                                   (set (re-seq #"\d+" my-number))))]
                     (reduce
                      (fn [cards match]
                        (update cards match (fn [n] (+ (get cards card-number) (or n 0)))))
                      cards
                      (range (inc card-number) (inc (+ matches card-number))))))
                 init-cards)
         vals
         (reduce +))))
