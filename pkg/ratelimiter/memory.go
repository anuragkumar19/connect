package ratelimiter

// type memoryStore struct {
// 	mu    sync.Mutex
// 	store map[string]*BucketCtx
// }

// func NewMemoryStore() *memoryStore {
// 	return &memoryStore{
// 		store: make(map[string]*BucketCtx),
// 		mu:    sync.Mutex{},
// 	}
// }

// var _ Store = (*memoryStore)(nil)

// func (s *memoryStore) Get(ctx context.Context, id string) (BucketCtx, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	v, ok := s.store[id]
// 	if !ok {
// 		return BucketCtx{}, ErrNotFound
// 	}
// 	return *v, nil
// }

// func (s *memoryStore) Create(ctx context.Context, bucket BucketCtx) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	_, ok := s.store[bucket.ID]
// 	if ok {
// 		return ErrAlreadyExist
// 	}

// 	s.store[bucket.ID] = &bucket
// 	return nil
// }

// func (s *memoryStore) Update(ctx context.Context, bucket BucketCtx) (BucketCtx, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	v, ok := s.store[bucket.ID]
// 	if !ok {
// 		return BucketCtx{}, ErrNotFound
// 	}

// 	if v.Revision != bucket.Revision {
// 		return BucketCtx{}, ErrRevisionMismatch
// 	}

// 	bucket.Revision++
// 	s.store[bucket.ID] = &bucket
// 	return bucket, nil
// }
